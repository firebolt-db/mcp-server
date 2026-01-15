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
	DatabasesFunc func(ctx context.Context) (*mcp.ReadResourceResult, error)
}

func (m *MockResourceFetcher) FetchCoreDatabaseResources(ctx context.Context) (*mcp.ReadResourceResult, error) {
	return m.DatabasesFunc(ctx)
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

func TestNewConnectCore(t *testing.T) {
	mock := &MockResourceFetcher{}
	connectTool := tool_connect_core.NewConnectCore(mock, nil, false)
	assert.NotNil(t, connectTool)
}

func TestConnectCore_Tool(t *testing.T) {
	mock := &MockResourceFetcher{}
	connectTool := tool_connect_core.NewConnectCore(mock, nil, false)

	tool := connectTool.Tool()
	assert.Equal(t, "firebolt_connect_core", tool.Name)
	assert.Contains(t, tool.Description, "Returns a list of Firebolt databases")
}

func TestConnectCore_Handler_Success_RequireProof(t *testing.T) {
	// Create test data
	databases := []string{"db1", "db2"}

	// Create mock fetcher
	mock := &MockResourceFetcher{
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, db := range databases {
				resources = append(resources, createDatabaseResource(db))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
	}

	// Create the tool
	connectCoreTool := tool_connect_core.NewConnectCore(mock, &validProof, false)

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

	// Check if we got the expected number of resources
	assert.Len(t, out.Results, len(databases))

	// Verify the content contains all expected resources
	resourceMap := make(map[string]bool)
	for _, content := range out.Results {
		embeddedResource, ok := content.(*mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		resourceMap[embeddedResource.Resource.URI] = true
	}

	// Check if all databases are present
	for _, db := range databases {
		uri := "firebolt://databases/" + db
		assert.True(t, resourceMap[uri], "Missing database resource: "+uri)
	}
}

func TestConnectCore_Handler_Success_NoProofRequired(t *testing.T) {
	// Create test data
	databases := []string{"db1", "db2"}
	// Create mock fetcher
	mock := &MockResourceFetcher{
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, db := range databases {
				resources = append(resources, createDatabaseResource(db))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
	}

	// Create the tool
	connectCoreTool := tool_connect_core.NewConnectCore(mock, nil, false)

	// Execute the handler
	request := &mcp.CallToolRequest{}
	in := tool_connect_core.Input{}
	result, out, err := connectCoreTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	assert.False(t, result.IsError)

	// Check if we got the expected number of resources
	assert.Len(t, out.Results, len(databases))

	// Verify the content contains all expected resources
	resourceMap := make(map[string]bool)
	for _, content := range out.Results {
		embeddedResource, ok := content.(*mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		resourceMap[embeddedResource.Resource.URI] = true
	}

	for _, db := range databases {
		uri := "firebolt://databases/" + db
		assert.True(t, resourceMap[uri], "Missing database resource: "+uri)
	}
}

func TestConnectCore_Handler_DatabasesFetchFailure(t *testing.T) {
	mock := &MockResourceFetcher{
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			return nil, errors.New("failed to fetch databases")
		},
	}

	connectCoreTool := tool_connect_core.NewConnectCore(mock, &validProof, false)
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

func TestConnectCore_Handler_DisableResources(t *testing.T) {
	// Create test data
	databases := []string{"db1"}

	// Create mock fetcher
	mock := &MockResourceFetcher{
		DatabasesFunc: func(ctx context.Context) (*mcp.ReadResourceResult, error) {
			var resources []*mcp.ResourceContents
			for _, db := range databases {
				resources = append(resources, createDatabaseResource(db))
			}
			return &mcp.ReadResourceResult{
				Contents: resources,
			}, nil
		},
	}

	// Create the tool with disableResources set to true
	connectCoreTool := tool_connect_core.NewConnectCore(mock, &validProof, true)

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

	// Check if we got the expected number of resources
	assert.Len(t, out.Results, len(databases))

	// Verify the content contains text content instead of embedded resources
	for _, content := range result.Content {
		textContent, ok := content.(*mcp.TextContent)
		require.True(t, ok, "Expected TextContent when disableResources is true")
		assert.NotEmpty(t, textContent.Text)
	}
}
