package resources_test

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
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
)

func TestDatabaseURI(t *testing.T) {
	tests := []struct {
		name     string
		account  string
		database string
		expected string
	}{
		{
			name:     "basic case",
			account:  "test-account",
			database: "test-database",
			expected: "firebolt://accounts/test-account/databases/test-database",
		},
		{
			name:     "with special characters",
			account:  "test.account",
			database: "test-database-123",
			expected: "firebolt://accounts/test.account/databases/test-database-123",
		},
		{
			name:     "empty values",
			account:  "",
			database: "",
			expected: "firebolt://accounts//databases/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resources.DatabaseURI(tt.account, tt.database)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewDatabases(t *testing.T) {
	pool := databasemock.NewPoolMock()
	databases := resources.NewDatabases(pool)
	assert.NotNil(t, databases)
}

func TestDatabases_ResourceTemplate(t *testing.T) {
	pool := databasemock.NewPoolMock()
	databases := resources.NewDatabases(pool)
	assert.NotEmpty(t, databases.ResourceTemplate())
}

func TestDatabases_Handler(t *testing.T) {
	tests := []struct {
		name          string
		params        map[string]any
		mockSetup     func(*databasemock.PoolMock, *databasemock.ConnectionMock)
		expected      []map[string]any
		expectedError string
	}{
		{
			name: "successful request",
			params: map[string]any{
				"account":  "test-account",
				"database": "test-database",
			},
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT database_name, description FROM information_schema.databases WHERE database_name = ?", sql)
					assert.Equal(t, []any{"test-database"}, args)
					return []map[string]any{
						{
							"database_name": "test-database",
							"description":   "Test database description",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expected: []map[string]any{
				{
					"account_name":  "test-account",
					"database_name": "test-database",
					"description":   "Test database description",
				},
			},
		},
		{
			name: "missing account parameter",
			params: map[string]any{
				"database": "test-database",
			},
			expectedError: "bad request: required argument account not provided",
		},
		{
			name: "missing database parameter",
			params: map[string]any{
				"account": "test-account",
			},
			expectedError: "bad request: required argument database not provided",
		},
		{
			name: "database connection error",
			params: map[string]any{
				"account":  "test-account",
				"database": "test-database",
			},
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				pool.WithGetConnectionFunc(func(params database.PoolParams) (database.Connection, error) {
					return nil, errors.New("connection error")
				})
			},
			expectedError: "failed to acquire database connection: connection error",
		},
		{
			name: "query execution error",
			params: map[string]any{
				"account":  "test-account",
				"database": "test-database",
			},
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					return nil, errors.New("query error")
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expectedError: "failed to query database: query error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := databasemock.NewPoolMock()
			conn := databasemock.NewConnectionMock()

			if tt.mockSetup != nil {
				tt.mockSetup(pool, conn)
			}

			databases := resources.NewDatabases(pool)

			request := mcp.ReadResourceRequest{}
			request.Params.Arguments = tt.params

			result, err := databases.Handler(t.Context(), request)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			require.Len(t, result, len(tt.expected))

			for i, res := range result {
				textRes, ok := res.(mcp.TextResourceContents)
				require.True(t, ok, "Expected TextResourceContents")

				var data map[string]any
				err := json.Unmarshal([]byte(textRes.Text), &data)
				require.NoError(t, err)

				assert.Equal(t, tt.expected[i], data)
				assert.Equal(t, resources.DatabaseURI("test-account", data["database_name"].(string)), textRes.URI)
				assert.Equal(t, mimetype.JSON, textRes.MIMEType)
			}
		})
	}
}

func TestDatabases_FetchDatabaseResources(t *testing.T) {
	tests := []struct {
		name          string
		account       string
		database      string
		mockSetup     func(*databasemock.PoolMock, *databasemock.ConnectionMock)
		expected      []map[string]any
		expectedError string
	}{
		{
			name:     "fetch specific database",
			account:  "test-account",
			database: "test-database",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT database_name, description FROM information_schema.databases WHERE database_name = ?", sql)
					assert.Equal(t, []any{"test-database"}, args)
					return []map[string]any{
						{
							"database_name": "test-database",
							"description":   "Production database",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expected: []map[string]any{
				{
					"account_name":  "test-account",
					"database_name": "test-database",
					"description":   "Production database",
				},
			},
		},
		{
			name:     "fetch all databases",
			account:  "test-account",
			database: "",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT database_name, description FROM information_schema.databases", sql)
					assert.Empty(t, args)
					return []map[string]any{
						{
							"database_name": "database-1",
							"description":   "First database",
						},
						{
							"database_name": "database-2",
							"description":   "Second database",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expected: []map[string]any{
				{
					"account_name":  "test-account",
					"database_name": "database-1",
					"description":   "First database",
				},
				{
					"account_name":  "test-account",
					"database_name": "database-2",
					"description":   "Second database",
				},
			},
		},
		{
			name:     "json marshaling error",
			account:  "test-account",
			database: "test-database",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					return []map[string]any{
						{
							"database_name": make(chan int), // Channels are not JSON-serializable
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expectedError: "failed to marshal row data to JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool := databasemock.NewPoolMock()
			conn := databasemock.NewConnectionMock()

			if tt.mockSetup != nil {
				tt.mockSetup(pool, conn)
			}

			databases := resources.NewDatabases(pool)

			result, err := databases.FetchDatabaseResources(t.Context(), tt.account, tt.database)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			require.Len(t, result, len(tt.expected))

			for i, res := range result {
				textRes, ok := res.(mcp.TextResourceContents)
				require.True(t, ok, "Expected TextResourceContents")

				var data map[string]any
				err := json.Unmarshal([]byte(textRes.Text), &data)
				require.NoError(t, err)

				assert.Equal(t, tt.expected[i], data)
				assert.Equal(t, resources.DatabaseURI(tt.account, data["database_name"].(string)), textRes.URI)
				assert.Equal(t, mimetype.JSON, textRes.MIMEType)
			}
		})
	}
}
