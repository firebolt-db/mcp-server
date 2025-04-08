package database_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"testing"

	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
)

func TestNewConnection(t *testing.T) {

	t.Run("successful connection", func(t *testing.T) {
		conn, closer, err := database.NewConnection(slogt.New(t), mockParams{})
		require.NoError(t, err)
		require.NotNil(t, conn)
		require.NotNil(t, closer)
		closer()
	})
}

func TestConnectionQuery(t *testing.T) {

	t.Run("successful query", func(t *testing.T) {

		conn, closer, err := database.NewConnection(slogt.New(t), mockParams{})
		require.NoError(t, err)
		t.Cleanup(closer)

		results, err := conn.Query(t.Context(), "SELECT col1, col2 FROM test")
		require.NoError(t, err)
		require.Len(t, results, 2)

		// Verify first row
		assert.Equal(t, "value1", results[0]["col1"])
		assert.Equal(t, 1, results[0]["col2"])

		// Verify second row
		assert.Equal(t, "value2", results[1]["col1"])
		assert.Equal(t, 2, results[1]["col2"])
	})

	t.Run("empty result set", func(t *testing.T) {

		conn, closer, err := database.NewConnection(slogt.New(t), mockParams{})
		require.NoError(t, err)
		t.Cleanup(closer)

		results, err := conn.Query(t.Context(), "SELECT * FROM empty")
		require.NoError(t, err)
		assert.Empty(t, results)
	})

	t.Run("query error", func(t *testing.T) {

		conn, closer, err := database.NewConnection(slogt.New(t), mockParams{})
		require.NoError(t, err)
		t.Cleanup(closer)

		results, err := conn.Query(t.Context(), "ERROR")
		require.Error(t, err)
		assert.Nil(t, results)
	})

	t.Run("closed connection", func(t *testing.T) {

		conn, closer, err := database.NewConnection(slogt.New(t), mockParams{})
		require.NoError(t, err)

		// Close the connection
		closer()

		// Try to query on a closed connection
		results, err := conn.Query(t.Context(), "SELECT col1, col2 FROM test")
		require.Error(t, err)
		assert.ErrorIs(t, err, database.ErrConnectionClosed)
		assert.Nil(t, results)
	})
}

func TestConnectionClose(t *testing.T) {

	t.Run("multiple close calls", func(t *testing.T) {

		conn, closer, err := database.NewConnection(slogt.New(t), mockParams{})
		require.NoError(t, err)

		// First close
		closer()

		// Second close should not panic
		closer()

		// Should still return error for operations
		results, err := conn.Query(t.Context(), "SELECT col1, col2 FROM test")
		require.Error(t, err)
		assert.ErrorIs(t, err, database.ErrConnectionClosed)
		assert.Nil(t, results)
	})
}

// Mock SQL driver implementation for testing
type mockDriver struct{}

type mockConn struct {
	closed bool
}

type mockStmt struct {
	query  string
	closed bool
}

type mockRows struct {
	columns []string
	data    [][]driver.Value
	closed  bool
	rowIdx  int
}

// Register the mock driver
func init() {
	sql.Register("fireboltmock", &mockDriver{})
}

func (d *mockDriver) Open(name string) (driver.Conn, error) {
	return &mockConn{}, nil
}

func (c *mockConn) Prepare(query string) (driver.Stmt, error) {
	if c.closed {
		return nil, driver.ErrBadConn
	}
	return &mockStmt{query: query}, nil
}

func (c *mockConn) Close() error {
	c.closed = true
	return nil
}

func (c *mockConn) Begin() (driver.Tx, error) {
	return nil, errors.New("transactions not supported")
}

func (s *mockStmt) Close() error {
	s.closed = true
	return nil
}

func (s *mockStmt) NumInput() int {
	return 0
}

func (s *mockStmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("exec not implemented")
}

func (s *mockStmt) Query(args []driver.Value) (driver.Rows, error) {

	if s.query == "ERROR" {
		return nil, errors.New("query error")
	}

	if s.query == "SELECT col1, col2 FROM test" {
		return &mockRows{
			columns: []string{"col1", "col2"},
			data: [][]driver.Value{
				{"value1", 1},
				{"value2", 2},
			},
		}, nil
	}

	return &mockRows{
		columns: []string{},
		data:    [][]driver.Value{},
	}, nil
}

func (r *mockRows) Columns() []string {
	return r.columns
}

func (r *mockRows) Close() error {
	r.closed = true
	return nil
}

func (r *mockRows) Next(dest []driver.Value) error {
	if r.rowIdx >= len(r.data) {
		return io.EOF
	}

	copy(dest, r.data[r.rowIdx])
	r.rowIdx++
	return nil
}

type mockParams struct {
}

func (m mockParams) DriverName() string {
	return "fireboltmock"
}

func (m mockParams) DSN() string {
	return "fireboltmock://test"
}
