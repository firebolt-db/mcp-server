package tool_docs_test

import (
	"context"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_docs"
)

func TestNewDocs(t *testing.T) {
	mock := &MockDocsFetcher{}
	docsTool := tool_docs.NewDocs(mock, false)
	assert.NotNil(t, docsTool)
}

func TestDocs_Tool(t *testing.T) {
	mock := &MockDocsFetcher{}
	docsTool := tool_docs.NewDocs(mock, false)

	tool := docsTool.Tool()
	assert.Equal(t, "firebolt_docs_overview", tool.Name)
	assert.Contains(t, tool.Description, "Returns Firebolt documentation overview")
}

func TestDocs_Handler_FetchError(t *testing.T) {
	// Create mock fetcher that returns an error
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) (*mcp.ReadResourceResult, error) {
			return nil, errors.New("failed to fetch article")
		},
	}

	// Create the tool
	docsTool := tool_docs.NewDocs(mock, false)

	// Execute the handler
	request := &mcp.CallToolRequest{}
	in := tool_docs.Input{}
	result, out, err := docsTool.Handler()(t.Context(), request, in)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover resources")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestDocs_Handler_DisableResources(t *testing.T) {
	// Create test data for default articles
	mockArticles := map[string]string{
		resources.DocsArticleOverview: "# Firebolt Overview\nThis is an overview of Firebolt.",
		resources.DocsArticleProof:    "# Proof Document\nSecret proof: proof_value_123",
		// resources.DocsArticleReference: "# Reference\nThis is the reference documentation.",
	}

	// Create mock fetcher that returns the mock articles
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) (*mcp.ReadResourceResult, error) {
			content, exists := mockArticles[article]
			if !exists {
				return nil, errors.New("article not found")
			}
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{createDocResource(article, content)},
			}, nil
		},
	}

	// Create the tool with disableResources set to true
	docsTool := tool_docs.NewDocs(mock, true)

	// Execute the handler with empty request (should return default articles)
	request := &mcp.CallToolRequest{}
	in := tool_docs.Input{}
	result, out, err := docsTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Should return 3 default articles
	require.NotNil(t, out)
	assert.Len(t, out.Articles, 2)

	// Verify the content contains text content instead of embedded resources
	textContents := make(map[string]string)
	for _, content := range out.Articles {
		textContent, ok := content.(*mcp.TextContent)
		require.True(t, ok, "Expected TextContent when disableResources is true")
		assert.NotEmpty(t, textContent.Text)

		// Store the text content for verification
		for articleID, expectedContent := range mockArticles {
			if textContent.Text == expectedContent {
				textContents[articleID] = textContent.Text
			}
		}
	}

	// Check if all default articles are present
	for articleID, expectedContent := range mockArticles {
		assert.Contains(t, textContents, articleID)
		assert.Equal(t, expectedContent, textContents[articleID])
	}
}

type MockDocsFetcher struct {
	FetchDocsFunc func(ctx context.Context, article string) (*mcp.ReadResourceResult, error)
}

func (m *MockDocsFetcher) FetchDocsResources(ctx context.Context, article string) (*mcp.ReadResourceResult, error) {
	return m.FetchDocsFunc(ctx, article)
}

// Helper to create a doc resource
func createDocResource(articleID, content string) *mcp.ResourceContents {
	return &mcp.ResourceContents{
		URI:      "firebolt://docs/" + articleID,
		MIMEType: mimetype.Markdown,
		Text:     content,
	}
}
