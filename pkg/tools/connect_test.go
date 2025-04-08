package tools_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

const validProof = "valid_proof"

// MockResourceFetcher is a test implementation of the resource fetcher interfaces
type MockResourceFetcher struct {
	AccountsFunc  func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error)
	DatabasesFunc func(ctx context.Context, accountName, databaseName string) ([]mcp.ResourceContents, error)
	EnginesFunc   func(ctx context.Context, accountName, engineName string) ([]mcp.ResourceContents, error)
}

func (m *MockResourceFetcher) FetchAccountResources(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
	return m.AccountsFunc(ctx, accountName)
}

func (m *MockResourceFetcher) FetchDatabaseResources(ctx context.Context, accountName, databaseName string) ([]mcp.ResourceContents, error) {
	return m.DatabasesFunc(ctx, accountName, databaseName)
}

func (m *MockResourceFetcher) FetchEngineResources(ctx context.Context, accountName, engineName string) ([]mcp.ResourceContents, error) {
	return m.EnginesFunc(ctx, accountName, engineName)
}

// Helpers to create resource mocks
func createAccountResource(name string) mcp.ResourceContents {
	account := map[string]string{
		"name":   name,
		"region": "us-east-1",
	}
	data, _ := json.Marshal(account)
	return mcp.TextResourceContents{
		URI:      "firebolt://accounts/" + name,
		MIMEType: mimetype.JSON,
		Text:     string(data),
	}
}

func createDatabaseResource(accountName, databaseName string) mcp.ResourceContents {
	database := map[string]string{
		"account_name":  accountName,
		"database_name": databaseName,
		"description":   "Description for " + databaseName,
	}
	data, _ := json.Marshal(database)
	return mcp.TextResourceContents{
		URI:      "firebolt://accounts/" + accountName + "/databases/" + databaseName,
		MIMEType: mimetype.JSON,
		Text:     string(data),
	}
}

func createEngineResource(accountName, engineName string) mcp.ResourceContents {
	engine := map[string]string{
		"account_name": accountName,
		"engine_name":  engineName,
		"status":       "running",
		"description":  "Description for " + engineName,
	}
	data, _ := json.Marshal(engine)
	return mcp.TextResourceContents{
		URI:      "firebolt://accounts/" + accountName + "/engines/" + engineName,
		MIMEType: mimetype.JSON,
		Text:     string(data),
	}
}

func TestNewConnect(t *testing.T) {
	mock := &MockResourceFetcher{}
	connectTool := tools.NewConnect(mock, mock, mock, validProof)
	assert.NotNil(t, connectTool)
}

func TestConnect_Tool(t *testing.T) {
	mock := &MockResourceFetcher{}
	connectTool := tools.NewConnect(mock, mock, mock, validProof)

	tool := connectTool.Tool()
	assert.Equal(t, "firebolt_connect", tool.Name)
	assert.Contains(t, tool.Description, "Returns a list of Firebolt accounts")
}

func TestConnect_Handler_Success(t *testing.T) {
	// Create test data
	accounts := []string{"account1", "account2"}
	databases := map[string][]string{
		"account1": {"db1", "db2"},
		"account2": {"db3"},
	}
	engines := map[string][]string{
		"account1": {"engine1"},
		"account2": {"engine2", "engine3"},
	}

	// Create mock fetcher
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
			var resources []mcp.ResourceContents
			for _, acc := range accounts {
				resources = append(resources, createAccountResource(acc))
			}
			return resources, nil
		},
		DatabasesFunc: func(ctx context.Context, accountName, databaseName string) ([]mcp.ResourceContents, error) {
			var resources []mcp.ResourceContents
			for _, db := range databases[accountName] {
				resources = append(resources, createDatabaseResource(accountName, db))
			}
			return resources, nil
		},
		EnginesFunc: func(ctx context.Context, accountName, engineName string) ([]mcp.ResourceContents, error) {
			var resources []mcp.ResourceContents
			for _, eng := range engines[accountName] {
				resources = append(resources, createEngineResource(accountName, eng))
			}
			return resources, nil
		},
	}

	// Create the tool
	connectTool := tools.NewConnect(mock, mock, mock, validProof)

	// Execute the handler
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"docs_proof": validProof,
	}
	result, err := connectTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Calculate expected total resources
	expectedCount := len(accounts) // accounts
	for _, dbs := range databases {
		expectedCount += len(dbs) // databases
	}
	for _, engs := range engines {
		expectedCount += len(engs) // engines
	}

	// Check if we got the expected number of resources
	assert.Len(t, result.Content, expectedCount)

	// Verify the content contains all expected resources
	resourceMap := make(map[string]bool)
	for _, content := range result.Content {
		embeddedResource, ok := content.(mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		textResource, ok := embeddedResource.Resource.(mcp.TextResourceContents)
		require.True(t, ok, "Expected TextResourceContents")

		resourceMap[textResource.URI] = true
	}

	// Check if all accounts are present
	for _, acc := range accounts {
		uri := "firebolt://accounts/" + acc
		assert.True(t, resourceMap[uri], "Missing account resource: "+uri)
	}

	// Check if all databases are present
	for acc, dbs := range databases {
		for _, db := range dbs {
			uri := "firebolt://accounts/" + acc + "/databases/" + db
			assert.True(t, resourceMap[uri], "Missing database resource: "+uri)
		}
	}

	// Check if all engines are present
	for acc, engs := range engines {
		for _, eng := range engs {
			uri := "firebolt://accounts/" + acc + "/engines/" + eng
			assert.True(t, resourceMap[uri], "Missing engine resource: "+uri)
		}
	}
}

func TestConnect_Handler_AccountFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
			return nil, errors.New("failed to fetch accounts")
		},
	}

	connectTool := tools.NewConnect(mock, mock, mock, validProof)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"docs_proof": validProof,
	}
	result, err := connectTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover resources")
	assert.Nil(t, result)
}

func TestConnect_Handler_InvalidAccountResource(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
			// Return a non-text resource to trigger the type assertion failure
			return []mcp.ResourceContents{
				mcp.BlobResourceContents{
					URI:      "firebolt://accounts/test-account",
					MIMEType: "application/octet-stream",
					Blob:     "binary data",
				},
			}, nil
		},
	}

	connectTool := tools.NewConnect(mock, mock, mock, validProof)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"docs_proof": validProof,
	}
	result, err := connectTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to convert account resource to text resource")
	assert.Nil(t, result)
}

func TestConnect_Handler_InvalidAccountJSON(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
			// Return invalid JSON to trigger unmarshal error
			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      "firebolt://accounts/test-account",
					MIMEType: mimetype.JSON,
					Text:     "invalid json",
				},
			}, nil
		},
	}

	connectTool := tools.NewConnect(mock, mock, mock, validProof)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"docs_proof": validProof,
	}
	result, err := connectTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal account resource")
	assert.Nil(t, result)
}

func TestConnect_Handler_DatabasesFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{createAccountResource("test-account")}, nil
		},
		DatabasesFunc: func(ctx context.Context, accountName, databaseName string) ([]mcp.ResourceContents, error) {
			return nil, errors.New("failed to fetch databases")
		},
	}

	connectTool := tools.NewConnect(mock, mock, mock, validProof)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"docs_proof": validProof,
	}
	result, err := connectTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover database resources")
	assert.Nil(t, result)
}

func TestConnect_Handler_EnginesFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{createAccountResource("test-account")}, nil
		},
		DatabasesFunc: func(ctx context.Context, accountName, databaseName string) ([]mcp.ResourceContents, error) {
			return []mcp.ResourceContents{}, nil
		},
		EnginesFunc: func(ctx context.Context, accountName, engineName string) ([]mcp.ResourceContents, error) {
			return nil, errors.New("failed to fetch engines")
		},
	}

	connectTool := tools.NewConnect(mock, mock, mock, validProof)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"docs_proof": validProof,
	}
	result, err := connectTool.Handler(t.Context(), request)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover engine resources")
	assert.Nil(t, result)
}
