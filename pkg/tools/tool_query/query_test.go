package tool_query_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/database/databasemock"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_query"
)

func TestNewQuery(t *testing.T) {
	mockPool := databasemock.NewPoolMock()
	queryTool := tool_query.NewQuery(mockPool)
	assert.NotNil(t, queryTool)
}

func TestQuery_Tool(t *testing.T) {
	mockPool := databasemock.NewPoolMock()
	queryTool := tool_query.NewQuery(mockPool)

	tool := queryTool.Tool()
	assert.Equal(t, "firebolt_query", tool.Name)
	assert.Contains(t, tool.Description, "Execute an SQL query against Firebolt")
}

func TestQuery_Handler_Success(t *testing.T) {

	// Create mock query result
	mockResult := []map[string]any{
		{
			"id":   1,
			"name": "test",
		},
		{
			"id":   2,
			"name": "test2",
		},
	}

	mockConnection := databasemock.NewConnectionMock().WithQueryFunc(
		func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
			assert.Equal(t, "SELECT * FROM test", sql)
			return mockResult, nil
		},
	)

	mockPool := databasemock.NewPoolMock()
	mockPool.RegisterConnection(database.PoolParams{
		AccountName:  "test-account",
		DatabaseName: ptrTo("test-db"),
		EngineName:   ptrTo("test-engine"),
	}, mockConnection)

	queryTool := tool_query.NewQuery(mockPool)

	// Create a request with all parameters
	request := &mcp.CallToolRequest{}
	in := tool_query.Input{
		Query:    "SELECT * FROM test",
		Account:  "test-account",
		Database: "test-db",
		Engine:   "test-engine",
	}

	// Execute the handler
	result, out, err := queryTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	assert.False(t, result.IsError)

	// Check the connection was requested with the right parameters
	callCount := mockPool.GetConnectionCallCount(database.PoolParams{
		AccountName:  "test-account",
		DatabaseName: ptrTo("test-db"),
		EngineName:   ptrTo("test-engine"),
	})
	assert.Equal(t, 1, callCount)

	// Verify the result contains the expected JSON data
	expectedJSON, _ := json.Marshal(mockResult)
	assert.Equal(t, string(expectedJSON), out.Result[0].(*mcp.TextContent).Text)
}

func TestQuery_Handler_MinimalParameters(t *testing.T) {
	// Create mock query result
	mockResult := []map[string]any{
		{
			"success": true,
		},
	}

	mockConnection := databasemock.NewConnectionMock().WithQueryFunc(
		func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
			assert.Equal(t, "SHOW ENGINES", sql)
			return mockResult, nil
		},
	)

	mockPool := databasemock.NewPoolMock()
	mockPool.RegisterConnection(database.PoolParams{
		AccountName: "test-account",
		// No database or engine specified
	}, mockConnection)

	queryTool := tool_query.NewQuery(mockPool)

	// Create a request with only required parameters
	request := &mcp.CallToolRequest{}
	in := tool_query.Input{
		Query:   "SHOW ENGINES",
		Account: "test-account",
	}

	// Execute the handler
	result, out, err := queryTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	assert.False(t, result.IsError)

	// Check the connection was requested with the right parameters
	callCount := mockPool.GetConnectionCallCount(database.PoolParams{
		AccountName: "test-account",
		// No database or engine
	})
	assert.Equal(t, 1, callCount)

	// Verify the result contains the expected JSON data
	expectedJSON, _ := json.Marshal(mockResult)
	assert.Equal(t, string(expectedJSON), out.Result[0].(*mcp.TextContent).Text)
}

func TestQuery_Handler_MissingRequiredParameters(t *testing.T) {
	mockPool := databasemock.NewPoolMock()
	queryTool := tool_query.NewQuery(mockPool)

	testCases := []struct {
		name      string
		in        tool_query.Input
		errSubstr string
	}{
		{
			name:      "missing query",
			in:        tool_query.Input{Account: "test-account"},
			errSubstr: "query",
		},
		{
			name:      "missing account",
			in:        tool_query.Input{Query: "SELECT 1"},
			errSubstr: "account",
		},
		{
			name:      "empty request",
			in:        tool_query.Input{},
			errSubstr: "bad request",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := &mcp.CallToolRequest{}
			in := tc.in

			result, out, err := queryTool.Handler()(t.Context(), request, in)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.errSubstr)
			assert.Nil(t, result)
			assert.Nil(t, out)
		})
	}
}

func TestQuery_Handler_ConnectionError(t *testing.T) {
	mockPool := databasemock.NewPoolMock().WithGetConnectionFunc(
		func(params database.PoolParams) (database.Connection, error) {
			return nil, errors.New("connection error")
		},
	)

	queryTool := tool_query.NewQuery(mockPool)

	request := &mcp.CallToolRequest{}
	in := tool_query.Input{
		Query:   "SELECT 1",
		Account: "test-account",
	}

	result, out, err := queryTool.Handler()(t.Context(), request, in)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to obtain database connection")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestQuery_Handler_QueryError(t *testing.T) {
	mockConnection := databasemock.NewConnectionMock().WithQueryFunc(
		func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
			return nil, errors.New("query execution error")
		},
	)

	mockPool := databasemock.NewPoolMock()
	mockPool.RegisterConnection(database.PoolParams{
		AccountName: "test-account",
	}, mockConnection)

	queryTool := tool_query.NewQuery(mockPool)

	request := &mcp.CallToolRequest{}
	in := tool_query.Input{
		Query:   "SELECT * FROM nonexistent_table",
		Account: "test-account",
	}

	result, out, err := queryTool.Handler()(t.Context(), request, in)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to query database")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestQuery_Handler_JSONMarshalError(t *testing.T) {
	// Create a value that can't be marshaled to JSON (a function)
	unmarshalableValue := func() {}
	mockResult := []map[string]any{
		{
			"unmarshalable": unmarshalableValue,
		},
	}

	mockConnection := databasemock.NewConnectionMock().WithQueryFunc(
		func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
			return mockResult, nil
		},
	)

	mockPool := databasemock.NewPoolMock()
	mockPool.RegisterConnection(database.PoolParams{
		AccountName: "test-account",
	}, mockConnection)

	queryTool := tool_query.NewQuery(mockPool)

	request := &mcp.CallToolRequest{}
	in := tool_query.Input{
		Query:   "SELECT problematic_data()",
		Account: "test-account",
	}

	result, out, err := queryTool.Handler()(t.Context(), request, in)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal query result")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func ptrTo[T any](v T) *T {
	return &v
}
