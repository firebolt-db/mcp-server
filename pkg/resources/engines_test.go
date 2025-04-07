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

func TestEngineURI(t *testing.T) {
	tests := []struct {
		name     string
		account  string
		engine   string
		expected string
	}{
		{
			name:     "basic case",
			account:  "test-account",
			engine:   "test-engine",
			expected: "firebolt://accounts/test-account/engines/test-engine",
		},
		{
			name:     "with special characters",
			account:  "test.account",
			engine:   "test-engine-123",
			expected: "firebolt://accounts/test.account/engines/test-engine-123",
		},
		{
			name:     "empty values",
			account:  "",
			engine:   "",
			expected: "firebolt://accounts//engines/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resources.EngineURI(tt.account, tt.engine)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNewEngines(t *testing.T) {
	pool := databasemock.NewPoolMock()
	engines := resources.NewEngines(pool)
	assert.NotNil(t, engines)
}

func TestEngines_ResourceTemplate(t *testing.T) {
	pool := databasemock.NewPoolMock()
	engines := resources.NewEngines(pool)
	assert.NotEmpty(t, engines.ResourceTemplate())
}

func TestEngines_Handler(t *testing.T) {
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
				"account": "test-account",
				"engine":  "test-engine",
			},
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT engine_name, description, status, version, type, family, nodes, clusters, auto_start FROM information_schema.engines WHERE engine_name = ?", sql)
					assert.Equal(t, []any{"test-engine"}, args)
					return []map[string]any{
						{
							"engine_name": "test-engine",
							"status":      "running",
							"version":     "3.0",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expected: []map[string]any{
				{
					"account_name": "test-account",
					"engine_name":  "test-engine",
					"status":       "running",
					"version":      "3.0",
				},
			},
		},
		{
			name: "missing account parameter",
			params: map[string]any{
				"engine": "test-engine",
			},
			expectedError: "bad request: required argument account not provided",
		},
		{
			name: "missing engine parameter",
			params: map[string]any{
				"account": "test-account",
			},
			expectedError: "bad request: required argument engine not provided",
		},
		{
			name: "database connection error",
			params: map[string]any{
				"account": "test-account",
				"engine":  "test-engine",
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
				"account": "test-account",
				"engine":  "test-engine",
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

			engines := resources.NewEngines(pool)

			request := mcp.ReadResourceRequest{}
			request.Params.Arguments = tt.params

			result, err := engines.Handler(t.Context(), request)

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
				assert.Equal(t, resources.EngineURI("test-account", data["engine_name"].(string)), textRes.URI)
				assert.Equal(t, mimetype.JSON, textRes.MIMEType)
			}
		})
	}
}

func TestEngines_FetchEngineResources(t *testing.T) {
	tests := []struct {
		name          string
		account       string
		engine        string
		mockSetup     func(*databasemock.PoolMock, *databasemock.ConnectionMock)
		expected      []map[string]any
		expectedError string
	}{
		{
			name:    "fetch specific engine",
			account: "test-account",
			engine:  "test-engine",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT engine_name, description, status, version, type, family, nodes, clusters, auto_start FROM information_schema.engines WHERE engine_name = ?", sql)
					assert.Equal(t, []any{"test-engine"}, args)
					return []map[string]any{
						{
							"engine_name": "test-engine",
							"status":      "running",
							"version":     "3.0",
							"type":        "general_purpose",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expected: []map[string]any{
				{
					"account_name": "test-account",
					"engine_name":  "test-engine",
					"status":       "running",
					"version":      "3.0",
					"type":         "general_purpose",
				},
			},
		},
		{
			name:    "fetch all engines",
			account: "test-account",
			engine:  "",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT engine_name, description, status, version, type, family, nodes, clusters, auto_start FROM information_schema.engines", sql)
					assert.Empty(t, args)
					return []map[string]any{
						{
							"engine_name": "engine-1",
							"status":      "running",
						},
						{
							"engine_name": "engine-2",
							"status":      "stopped",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{AccountName: "test-account"}, conn)
			},
			expected: []map[string]any{
				{
					"account_name": "test-account",
					"engine_name":  "engine-1",
					"status":       "running",
				},
				{
					"account_name": "test-account",
					"engine_name":  "engine-2",
					"status":       "stopped",
				},
			},
		},
		{
			name:    "json marshaling error",
			account: "test-account",
			engine:  "test-engine",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					return []map[string]any{
						{
							"engine_name": make(chan int), // Channels are not JSON-serializable
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

			engines := resources.NewEngines(pool)

			result, err := engines.FetchEngineResources(t.Context(), tt.account, tt.engine)

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
				assert.Equal(t, resources.EngineURI(tt.account, data["engine_name"].(string)), textRes.URI)
				assert.Equal(t, mimetype.JSON, textRes.MIMEType)
			}
		})
	}
}
