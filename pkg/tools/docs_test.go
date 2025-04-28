package tools_test

import (
	"context"
	"errors"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

func TestNewDocs(t *testing.T) {
	mock := &MockDocsFetcher{}
	docsTool := tools.NewDocs(mock, false)
	assert.NotNil(t, docsTool)
}

func TestDocs_Tool(t *testing.T) {
	mock := &MockDocsFetcher{}
	docsTool := tools.NewDocs(mock, false)

	tool := docsTool.Tool()
	assert.Equal(t, "firebolt_docs", tool.Name)
	assert.Contains(t, tool.Description, "Returns Firebolt documentation articles")
}

func TestDocs_Handler_DefaultArticles(t *testing.T) {
	// Create test data for default articles
	mockArticles := map[string]string{
		resources.DocsArticleOverview:  "# Firebolt Overview\nThis is an overview of Firebolt.",
		resources.DocsArticleProof:     "# Proof Document\nSecret proof: proof_value_123",
		resources.DocsArticleReference: "# Reference\nThis is the reference documentation.",
	}

	// Create mock fetcher that returns the mock articles
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) ([]mcp.ResourceContents, error) {
			content, exists := mockArticles[article]
			if !exists {
				return nil, errors.New("article not found")
			}
			return []mcp.ResourceContents{createDocResource(article, content)}, nil
		},
	}

	// Create the tool
	docsTool := tools.NewDocs(mock, false)

	// Execute the handler with empty request (should return default articles)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{}
	result, err := docsTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Should return 3 default articles
	assert.Len(t, result.Content, 3)

	// Verify the content contains all expected resources
	resourceMap := make(map[string]string)
	for _, content := range result.Content {
		embeddedResource, ok := content.(mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		textResource, ok := embeddedResource.Resource.(mcp.TextResourceContents)
		require.True(t, ok, "Expected TextResourceContents")

		resourceMap[textResource.URI] = textResource.Text
	}

	// Check if all default articles are present
	for articleID, expectedContent := range mockArticles {
		uri := "firebolt://docs/" + articleID
		assert.Contains(t, resourceMap, uri)
		assert.Equal(t, expectedContent, resourceMap[uri])
	}
}

func TestDocs_Handler_SpecificArticles(t *testing.T) {
	// Create test data for specific articles
	mockArticles := map[string]string{
		"article1": "# Article 1\nContent for article 1",
		"article2": "# Article 2\nContent for article 2",
	}

	// Create mock fetcher that returns the mock articles
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) ([]mcp.ResourceContents, error) {
			content, exists := mockArticles[article]
			if !exists {
				return nil, errors.New("article not found")
			}
			return []mcp.ResourceContents{createDocResource(article, content)}, nil
		},
	}

	// Create the tool
	docsTool := tools.NewDocs(mock, false)

	// Execute the handler with specific articles
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"articles": []any{"article1", "article2"},
	}
	result, err := docsTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Should return the 2 requested articles
	assert.Len(t, result.Content, 2)

	// Verify the content contains all expected resources
	resourceMap := make(map[string]string)
	for _, content := range result.Content {
		embeddedResource, ok := content.(mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		textResource, ok := embeddedResource.Resource.(mcp.TextResourceContents)
		require.True(t, ok, "Expected TextResourceContents")

		resourceMap[textResource.URI] = textResource.Text
	}

	// Check if all requested articles are present
	for articleID, expectedContent := range mockArticles {
		uri := "firebolt://docs/" + articleID
		assert.Contains(t, resourceMap, uri)
		assert.Equal(t, expectedContent, resourceMap[uri])
	}
}

func TestDocs_Handler_FetchError(t *testing.T) {
	// Create mock fetcher that returns an error
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) ([]mcp.ResourceContents, error) {
			return nil, errors.New("failed to fetch article")
		},
	}

	// Create the tool
	docsTool := tools.NewDocs(mock, false)

	// Execute the handler
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{}
	result, err := docsTool.Handler(t.Context(), request)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to discover resources")
	assert.Nil(t, result)
}

func TestDocs_Handler_InvalidArticleID(t *testing.T) {
	// Create the tool with any mock
	mock := &MockDocsFetcher{}
	docsTool := tools.NewDocs(mock, false)

	// Execute the handler with an invalid article ID type
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"articles": []any{123}, // Not a string
	}
	result, err := docsTool.Handler(t.Context(), request)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid type for article ID")
	assert.Nil(t, result)
}

func TestDocs_Handler_MultipleFetchedResources(t *testing.T) {
	// Create mock fetcher that returns multiple resources for a single article ID
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) ([]mcp.ResourceContents, error) {
			if article == "multi-resource" {
				return []mcp.ResourceContents{
					createDocResource("multi-resource-1", "# Part 1\nContent part 1"),
					createDocResource("multi-resource-2", "# Part 2\nContent part 2"),
				}, nil
			}
			return nil, errors.New("article not found")
		},
	}

	// Create the tool
	docsTool := tools.NewDocs(mock, false)

	// Execute the handler
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{
		"articles": []any{"multi-resource"},
	}
	result, err := docsTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Should return both resources
	assert.Len(t, result.Content, 2)

	// Verify both parts are present
	resourceMap := make(map[string]bool)
	for _, content := range result.Content {
		embeddedResource, ok := content.(mcp.EmbeddedResource)
		require.True(t, ok, "Expected EmbeddedResource")

		textResource, ok := embeddedResource.Resource.(mcp.TextResourceContents)
		require.True(t, ok, "Expected TextResourceContents")

		resourceMap[textResource.URI] = true
	}

	assert.True(t, resourceMap["firebolt://docs/multi-resource-1"])
	assert.True(t, resourceMap["firebolt://docs/multi-resource-2"])
}

func TestDocs_Handler_DisableResources(t *testing.T) {
	// Create test data for default articles
	mockArticles := map[string]string{
		resources.DocsArticleOverview:  "# Firebolt Overview\nThis is an overview of Firebolt.",
		resources.DocsArticleProof:     "# Proof Document\nSecret proof: proof_value_123",
		resources.DocsArticleReference: "# Reference\nThis is the reference documentation.",
	}

	// Create mock fetcher that returns the mock articles
	mock := &MockDocsFetcher{
		FetchDocsFunc: func(ctx context.Context, article string) ([]mcp.ResourceContents, error) {
			content, exists := mockArticles[article]
			if !exists {
				return nil, errors.New("article not found")
			}
			return []mcp.ResourceContents{createDocResource(article, content)}, nil
		},
	}

	// Create the tool with disableResources set to true
	docsTool := tools.NewDocs(mock, true)

	// Execute the handler with empty request (should return default articles)
	request := mcp.CallToolRequest{}
	request.Params.Arguments = map[string]any{}
	result, err := docsTool.Handler(t.Context(), request)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	// Should return 3 default articles
	assert.Len(t, result.Content, 3)

	// Verify the content contains text content instead of embedded resources
	textContents := make(map[string]string)
	for _, content := range result.Content {
		textContent, ok := content.(mcp.TextContent)
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
	FetchDocsFunc func(ctx context.Context, article string) ([]mcp.ResourceContents, error)
}

func (m *MockDocsFetcher) FetchDocsResources(ctx context.Context, article string) ([]mcp.ResourceContents, error) {
	return m.FetchDocsFunc(ctx, article)
}

// Helper to create a doc resource
func createDocResource(articleID, content string) mcp.ResourceContents {
	return mcp.TextResourceContents{
		URI:      "firebolt://docs/" + articleID,
		MIMEType: mimetype.Markdown,
		Text:     content,
	}
}
