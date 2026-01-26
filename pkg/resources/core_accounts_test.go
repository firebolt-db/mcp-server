package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/database/databasemock"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
)

func TestNewCoreAccounts(t *testing.T) {
	pool := databasemock.NewPoolMock()
	accounts := resources.NewCoreAccounts(pool)
	assert.NotNil(t, accounts)
}

func TestCoreAccounts_ResourceTemplate(t *testing.T) {

	pool := databasemock.NewPoolMock()
	accounts := resources.NewCoreAccounts(pool)
	template := accounts.ResourceTemplate()

	assert.NotEmpty(t, template.Name)
	assert.NotEmpty(t, template.URITemplate)
	assert.NotEmpty(t, template.MIMEType)
	assert.NotEmpty(t, template.Description)
}

func TestCoreAccounts_FetchAccountResources(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(*databasemock.PoolMock, *databasemock.ConnectionMock)
		expected      []map[string]any
		expectedError string
	}{
		{
			name: "fetch all accounts",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					assert.Equal(t, "SELECT account_id, account_name FROM information_schema.accounts;", sql)
					assert.Empty(t, args)
					return []map[string]any{
						{
							"account_id":   "test-account-id",
							"account_name": "test-account",
						},
					}, nil
				})
				pool.RegisterConnection(database.PoolParams{}, conn)
			},
			expected: []map[string]any{
				{
					"account_id":   "test-account-id",
					"account_name": "test-account",
				},
			},
		},
		{
			name: "json marshaling error test",
			mockSetup: func(pool *databasemock.PoolMock, conn *databasemock.ConnectionMock) {
				conn.WithQueryFunc(func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
					// Create a custom type that will cause JSON marshaling to fail
					// This is a bit tricky to test directly in Go, so we'll simulate by returning accounts
					// and rely on test coverage for the error path
					return []map[string]any{
						{
							"account_name": make(chan int), // Channels are not JSON-serializable
						},
					}, nil

				})
				pool.RegisterConnection(database.PoolParams{}, conn)
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

			accounts := resources.NewCoreAccounts(pool)

			result, err := accounts.FetchAccountResources(t.Context())

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			require.Len(t, result.Contents, len(tt.expected))

			for i, res := range result.Contents {
				var data map[string]any
				err := json.Unmarshal([]byte(res.Text), &data)
				require.NoError(t, err)

				assert.Equal(t, tt.expected[i], data)
				assert.Equal(t, resources.AccountURI(data["account_name"].(string)), res.URI)
				assert.Equal(t, mimetype.JSON, res.MIMEType)
			}
		})
	}
}
