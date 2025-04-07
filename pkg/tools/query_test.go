package tools_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/database/databasemock"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

func TestNewQuery(t *testing.T) {
	mockPool := databasemock.NewPoolMock()
	queryTool := tools.NewQuery(mockPool)
	assert.NotNil(t, queryTool)
}

func TestQuery_Tool(t *testing.T) {
	mockPool := databasemock.NewPoolMock()
	queryTool := tools.NewQuery(mockPool)

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

	queryTool := tools.NewQuery(mockPool)

	// Create a request with all parameters
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"query":    "SELECT * FROM test",
		"account":  "test-account",
		"database": "test-db",
		"engine":   "test-engine",
	}

	// Execute the handler
	result, err := queryTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
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
	assert.Equal(t, string(expectedJSON), result.Content[0].(mcp.TextContent).Text)
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

	queryTool := tools.NewQuery(mockPool)

	// Create a request with only required parameters
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"query":   "SHOW ENGINES",
		"account": "test-account",
		// No database or engine
	}

	// Execute the handler
	result, err := queryTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Check the connection was requested with the right parameters
	callCount := mockPool.GetConnectionCallCount(database.PoolParams{
		AccountName: "test-account",
		// No database or engine
	})
	assert.Equal(t, 1, callCount)

	// Verify the result contains the expected JSON data
	expectedJSON, _ := json.Marshal(mockResult)
	assert.Equal(t, string(expectedJSON), result.Content[0].(mcp.TextContent).Text)
}

func TestQuery_Handler_MissingRequiredParameters(t *testing.T) {
	mockPool := databasemock.NewPoolMock()
	queryTool := tools.NewQuery(mockPool)

	testCases := []struct {
		name      string
		arguments map[string]any
		errSubstr string
	}{
		{
			name:      "missing query",
			arguments: map[string]any{"account": "test-account"},
			errSubstr: "query",
		},
		{
			name:      "missing account",
			arguments: map[string]any{"query": "SELECT 1"},
			errSubstr: "account",
		},
		{
			name:      "empty request",
			arguments: map[string]any{},
			errSubstr: "bad request",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request := mcp.CallToolRequest{}
			request.Params.Arguments = tc.arguments

			result, err := queryTool.Handler(t.Context(), request)

			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.errSubstr)
			assert.Nil(t, result)
		})
	}
}

func TestQuery_Handler_ConnectionError(t *testing.T) {
	mockPool := databasemock.NewPoolMock().WithGetConnectionFunc(
		func(params database.PoolParams) (database.Connection, error) {
			return nil, errors.New("connection error")
		},
	)

	queryTool := tools.NewQuery(mockPool)

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"query":   "SELECT 1",
		"account": "test-account",
	}

	result, err := queryTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to obtain database connection")
	assert.Nil(t, result)
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

	queryTool := tools.NewQuery(mockPool)

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"query":   "SELECT * FROM nonexistent_table",
		"account": "test-account",
	}

	result, err := queryTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to query database")
	assert.Nil(t, result)
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

	queryTool := tools.NewQuery(mockPool)

	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"query":   "SELECT problematic_data()",
		"account": "test-account",
	}

	result, err := queryTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal query result")
	assert.Nil(t, result)
}

func ptrTo[T any](v T) *T {
	return &v
}
