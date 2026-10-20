package tool_query_test

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/database/databasemock"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_query"
)

func TestQuery_Tool_Annotations(t *testing.T) {
	mockPool := databasemock.NewPoolMock()

	t.Run("default has no annotations", func(t *testing.T) {
		tool := tool_query.NewQuery(mockPool).Tool()
		assert.Nil(t, tool.Annotations)
		assert.Equal(t, "Execute an SQL query against Firebolt.", tool.Description)
	})

	t.Run("read-only declares read-only hints", func(t *testing.T) {
		tool := tool_query.NewQuery(mockPool, tool_query.WithReadOnly()).Tool()
		require.NotNil(t, tool.Annotations)
		assert.True(t, tool.Annotations.ReadOnlyHint)
		assert.True(t, tool.Annotations.IdempotentHint)
		require.NotNil(t, tool.Annotations.DestructiveHint)
		assert.False(t, *tool.Annotations.DestructiveHint)
		assert.Nil(t, tool.Annotations.OpenWorldHint, "SELECT can read external data such as S3 files")
		assert.Contains(t, tool.Description, "Read-only")
	})
}

func TestQuery_Handler_ReadOnly_Allows(t *testing.T) {
	queries := []string{
		"SELECT 1",
		"select * from t",
		"  \n\tSELECT 1",
		"SELECT 1;",
		"SELECT 1;  \n",
		"-- leading comment\nSELECT 1",
		"/* leading */ SELECT 1",
		"WITH x AS (SELECT 1 AS a) SELECT a FROM x",
		"SELECT 'drop table t; insert into t' AS s",
		"SELECT 'it''s; delete' AS s",
		`SELECT E'it\'s; delete' AS s`,
		`SELECT "update", "delete" FROM t`,
		"SELECT 1 -- DROP TABLE t",
		"SELECT updated_at, created_by FROM t",
		"SELECT $$drop table t -- not a comment$$ AS s",
		"SELECT a$b, $c, $d$ FROM t",
		"SELECT * FROM t WHERE a = $1",
		`SELECT regexp_like(s, '\d+') FROM t`,
	}

	for _, q := range queries {
		t.Run(q, func(t *testing.T) {
			var executed string
			conn := databasemock.NewConnectionMock().WithQueryFunc(
				func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					executed = sql
					return []map[string]any{{"ok": true}}, nil
				},
			)
			pool := databasemock.NewPoolMock()
			pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)

			queryTool := tool_query.NewQuery(pool, tool_query.WithReadOnly())
			_, out, err := queryTool.Handler()(t.Context(), &mcp.CallToolRequest{}, tool_query.Input{
				Query:   q,
				Account: "test-account",
			})

			require.NoError(t, err)
			require.NotNil(t, out)
			assert.Equal(t, q, executed)
		})
	}
}

func TestQuery_Handler_ReadOnly_Rejects(t *testing.T) {
	queries := []string{
		"INSERT INTO t VALUES (1)",
		"insert into t select * from s",
		"UPDATE t SET a = 1",
		"DELETE FROM t",
		"TRUNCATE TABLE t",
		"DROP TABLE t",
		"CREATE TABLE t (a INT)",
		"ALTER ENGINE e SET NODES = 2",
		"GRANT SELECT ON t TO r",
		"REVOKE SELECT ON t FROM r",
		"VACUUM t",
		"COPY t FROM 's3://bucket/'",
		"CALL recommend_ddl(t, (SELECT 1))",
		"STOP ENGINE e",
		"SHOW ENGINES",
		"/* x */ DROP TABLE t",
		"-- x\nDROP TABLE t",
		"SELECT 1; DROP TABLE t",
		"SELECT 1;DROP TABLE t",
		"SELECT 1; SELECT 2",
		"SELECT 1;;",
		"SELECT 1; 'x'",
		"WITH x AS (SELECT 1) INSERT INTO t SELECT * FROM x",
		"SELECT * FROM t; -- ok\nDELETE FROM t",
		`SELECT E'\'' ; DROP TABLE t; --'`,
		"SELECT $$ -- $$; DROP TABLE t",
		"SELECT $$ /* $$; DROP TABLE t /* */",
		"SELECT $$ ' $$; DELETE FROM t --'",
		"WITH x AS (SELECT $$--$$) INSERT INTO t SELECT * FROM x",
		"SELECT 1 AS $tag$; DROP TABLE t; $tag$",
		"SELECT $q$it's $$ nested; delete$q$ AS s",
		"SELECT 1$$ -- $$; DROP TABLE t",
		`SELECT 'x\' AS a -- ' ; DROP TABLE t`,
		`SELECT 1 FROM t WHERE c = 'C:\' -- ' ; DROP TABLE t`,
		`SELECT 1 AS "x\" -- " ; DROP TABLE t`,
		`SELECT 'x\' -- ' ; SET foo = 1`,
		"SELECT $$ ' $$ -- ' ; DROP TABLE t",
		"SELECT $$drop table t; -- not a comment$$ AS s",
		"SELECT 1; -- trailing comment",
		"SELECT 1; /* trailing */",
		`SELECT 1 AS e, 'a\' ; DROP TABLE t; --'`,
		"SELECT 1;\r\n",
		"SELECT 1 AS x INTO t_new",
		"SELECT 1 INTO TABLE t_new",
		"WITH x AS (SELECT 1 AS a) SELECT a INTO t_new FROM x",
		"SELECT $$unterminated",
		"SELECT 'unterminated",
		"SELECT /* unterminated",
		"",
		"   ",
		"-- only a comment",
	}

	for _, q := range queries {
		t.Run(q, func(t *testing.T) {
			pool := databasemock.NewPoolMock()
			queryTool := tool_query.NewQuery(pool, tool_query.WithReadOnly())

			result, out, err := queryTool.Handler()(t.Context(), &mcp.CallToolRequest{}, tool_query.Input{
				Query:   q,
				Account: "test-account",
			})

			require.Error(t, err)
			assert.Contains(t, err.Error(), "bad request")
			assert.Nil(t, result)
			assert.Nil(t, out)
			assert.Equal(t, 0, pool.GetConnectionCallCount(database.PoolParams{AccountName: "test-account"}),
				"a rejected query must not reach the database")
		})
	}
}

func TestQuery_Handler_Default_AllowsWrites(t *testing.T) {
	var executed string
	conn := databasemock.NewConnectionMock().WithQueryFunc(
		func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
			executed = sql
			return []map[string]any{}, nil
		},
	)
	pool := databasemock.NewPoolMock()
	pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)

	queryTool := tool_query.NewQuery(pool)
	_, _, err := queryTool.Handler()(t.Context(), &mcp.CallToolRequest{}, tool_query.Input{
		Query:   "INSERT INTO t VALUES (1)",
		Account: "test-account",
	})

	require.NoError(t, err)
	assert.Equal(t, "INSERT INTO t VALUES (1)", executed)
}
