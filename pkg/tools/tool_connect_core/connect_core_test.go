package tool_connect_core_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_connect_core"
)

var validProof = "valid_proof"

// MockResourceFetcher is a test implementation of the resource fetcher interfaces
type MockResourceFetcher struct {
	AccountsFunc  func(ctx context.Context) (*mcp.ReadResourceResult, error)
	DatabasesFunc func(ctx context.Context) (*mcp.ReadResourceResult, error)
	EnginesFunc   func(ctx context.Context) (*mcp.ReadResourceResult, error)
}

func (m *MockResourceFetcher) FetchAccountResources(ctx context.Context) (*mcp.ReadResourceResult, error) {
	return m.AccountsFunc(ctx)
}

func (m *MockResourceFetcher) FetchDatabaseResources(ctx context.Context, _, _ string) (*mcp.ReadResourceResult, error) {
	return m.DatabasesFunc(ctx)
}

func (m *MockResourceFetcher) FetchEngineResources(ctx context.Context, _, _ string) (*mcp.ReadResourceResult, error) {
	return m.EnginesFunc(ctx)
}

// Helpers to create resource mocks
func createAccountResource(name string) *mcp.ResourceContents {
	account := map[string]string{
		"name": name,
	}
	data, _ := json.Marshal(account)
	return &mcp.ResourceContents{
		URI:      "firebolt://accounts/" + name,
		MIMEType: mimetype.JSON,
		Text:     string(data),
	}
}
func createDatabaseResource(databaseName string) *mcp.ResourceContents {
	database := map[string]string{
		"database_name": databaseName,
		"description":   "Description for " + databaseName,
	}
	data, _ := json.Marshal(database)
	return &mcp.ResourceContents{
		URI:      "firebolt://databases/" + databaseName,
		MIMEType: mimetype.JSON,
		Text:     string(data),
	}
}

func createEngineResource(engineName string) *mcp.ResourceContents {
	engine := map[string]string{
		"engine_name": engineName,
		"status":      "running",
		"description": "Description for " + engineName,
	}
	data, _ := json.Marshal(engine)
	return &mcp.ResourceContents{
		URI:      "firebolt://engines/" + engineName,
		MIMEType: mimetype.JSON,
		Text:     string(data),
	}
}

func TestNewConnectCore(t *testing.T) {
	mock := &MockResourceFetcher{}
	connectTool := tool_connect_core.NewConnectCore(mock, mock, mock, nil, false)
	assert.NotNil(t, connectTool)
}

func TestConnectCore_Tool(t *testing.T) {
	mock := &MockResourceFetcher{}
	connectTool := tool_connect_core.NewConnectCore(mock, mock, mock, nil, false)

	tool := connectTool.Tool()
	assert.Equal(t, "firebolt_connect_core", tool.Name)
	assert.Contains(t, tool.Description, "Returns a list of databases you have access to in Firebolt Core")
}

func TestConnectCore_Handler_Success_RequireProof(t *testing.T) {
	// Create test data
	accounts := []string{"account1", "account2"}
	databases := []string{"db1", "db2"}
	engines := []string{"engine1", "engine2"}

	// Create mock fetcher
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, acc := range accounts {
				resources = append(resources, createAccountResource(acc))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, db := range databases {
				resources = append(resources, createDatabaseResource(db))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
		EnginesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, eng := range engines {
				resources = append(resources, createEngineResource(eng))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
	}

	// Create the tool
	connectCoreTool := tool_connect_core.NewConnectCore(mock, mock, mock, &validProof, false)

	// Execute the handler
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{
		DocsProof: validProof,
	}
	result, out, err := connectCoreTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	assert.False(t, result.IsError)

	expectedCount := len(accounts) + len(databases) + len(engines)

	// Check if we got the expected number of resources
	assert.Len(t, out.Results, expectedCount)

	// Verify the content contains all expected resources
	resourceMap := make(map[string]bool)
	for _, content := range out.Results {
		embeddedResource, ok := content.(*mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		resourceMap[embeddedResource.Resource.URI] = true
	}

	// Check if all databases are present
	for _, acc := range accounts {
		uri := "firebolt://accounts/" + acc
		assert.True(t, resourceMap[uri], "Missing account resource: "+uri)
	}

	for _, db := range databases {
		uri := "firebolt://databases/" + db
		assert.True(t, resourceMap[uri], "Missing database resource: "+uri)
	}

	for _, eng := range engines {
		uri := "firebolt://engines/" + eng
		assert.True(t, resourceMap[uri], "Missing engine resource: "+uri)
	}
}

func TestConnectCore_Handler_Success_NoProofRequired(t *testing.T) {
	// Create test data
	accounts := []string{"account1", "account2"}
	databases := []string{"db1", "db2"}
	engines := []string{"engine1", "engine2"}

	// Create mock fetcher
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, acc := range accounts {
				resources = append(resources, createAccountResource(acc))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, db := range databases {
				resources = append(resources, createDatabaseResource(db))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
		EnginesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, eng := range engines {
				resources = append(resources, createEngineResource(eng))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
	}

	// Create the tool
	connectCoreTool := tool_connect_core.NewConnectCore(mock, mock, mock, nil, false)

	// Execute the handler
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{}
	result, out, err := connectCoreTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	assert.False(t, result.IsError)

	expectedCount := len(accounts) + len(databases) + len(engines)

	// Check if we got the expected number of resources
	assert.Len(t, out.Results, expectedCount)

	// Verify the content contains all expected resources
	resourceMap := make(map[string]bool)
	for _, content := range out.Results {
		embeddedResource, ok := content.(*mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		resourceMap[embeddedResource.Resource.URI] = true
	}

	// Check if all databases are present
	for _, acc := range accounts {
		uri := "firebolt://accounts/" + acc
		assert.True(t, resourceMap[uri], "Missing account resource: "+uri)
	}

	for _, db := range databases {
		uri := "firebolt://databases/" + db
		assert.True(t, resourceMap[uri], "Missing database resource: "+uri)
	}

	for _, eng := range engines {
		uri := "firebolt://engines/" + eng
		assert.True(t, resourceMap[uri], "Missing engine resource: "+uri)
	}
}

func TestConnectCore_Handler_AccountFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return nil, errors.New("failed to fetch accounts")
		},
	}

	connectTool := tool_connect_core.NewConnectCore(mock, mock, mock, &validProof, false)
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{
		DocsProof: validProof,
	}
	result, out, err := connectTool.Handler()(t.Context(), request, in)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover account resources")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestConnectCore_Handler_DatabasesFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{createAccountResource("test-account")},
			}, nil
		},
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return nil, errors.New("failed to fetch databases")
		},
	}

	connectCoreTool := tool_connect_core.NewConnectCore(mock, mock, mock, &validProof, false)
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{
		DocsProof: validProof,
	}
	result, out, err := connectCoreTool.Handler()(t.Context(), request, in)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover database resources")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestConnectCore_Handler_EnginesFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{createAccountResource("test-account")},
			}, nil
		},
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{},
			}, nil
		},
		EnginesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return nil, errors.New("failed to fetch engines")
		},
	}

	connectTool := tool_connect_core.NewConnectCore(mock, mock, mock, &validProof, false)
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{
		DocsProof: validProof,
	}
	result, out, err := connectTool.Handler()(t.Context(), request, in)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover engine resources")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestConnectCore_Handler_DisableResources(t *testing.T) {
	// Create test data
	accounts := []string{"account1", "account2"}
	databases := []string{"db1", "db2"}
	engines := []string{"engine1", "engine2"}

	// Create mock fetcher
	mock := &MockResourceFetcher{
		AccountsFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, acc := range accounts {
				resources = append(resources, createAccountResource(acc))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, db := range databases {
				resources = append(resources, createDatabaseResource(db))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
		EnginesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, eng := range engines {
				resources = append(resources, createEngineResource(eng))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
	}

	// Create the tool with disableResources set to true
	connectCoreTool := tool_connect_core.NewConnectCore(mock, mock, mock, &validProof, true)

	// Execute the handler
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{
		DocsProof: validProof,
	}
	result, out, err := connectCoreTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	assert.False(t, result.IsError)

	expectedCount := len(accounts) + len(databases) + len(engines)

	// Check if we got the expected number of resources
	assert.Len(t, out.Results, expectedCount)

	// Verify the content contains text content instead of embedded resources
	for _, content := range out.Results {
		textContent, ok := content.(*mcp.TextContent)
		require.True(t, ok, "Expected TextContent when disableResources is true")
		assert.NotEmpty(t, textContent.Text)
	}
}
