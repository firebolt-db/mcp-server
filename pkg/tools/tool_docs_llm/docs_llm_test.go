package tool_docs_llm_test

import (
	"context"
	"errors"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_docs_llm"
)

func TestNewDocsLLM(t *testing.T) {
	mock := &MockDocsFetcher{}
	docsLLMTool := tool_docs_llm.NewDocsLLM(mock, false)
	assert.NotNil(t, docsLLMTool)
}

func TestDocsLLM_Tool(t *testing.T) {
	mock := &MockDocsFetcher{}
	docsLLMTool := tool_docs_llm.NewDocsLLM(mock, false)

	tool := docsLLMTool.Tool()
	assert.Equal(t, "firebolt_docs_llm", tool.Name)
	assert.Contains(t, tool.Description, "Returns Firebolt documentation articles prepared to be used by LLMs")
}

func TestDocsLLM_Handler_FetchError(t *testing.T) {
	// Create mock fetcher that returns an error
	mock := &MockDocsFetcher{
		FetchLLMDocsFunc: func(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error) {
			return nil, errors.New("failed to fetch article")
		},
	}

	// Create the tool
	docsLLMTool := tool_docs_llm.NewDocsLLM(mock, false)

	// Execute the handler
	request := &mcp.CallToolRequest{}
	in := tool_docs_llm.Input{}
	result, out, err := docsLLMTool.Handler()(t.Context(), request, in)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch LLM docs resource")
	assert.Nil(t, result)
	assert.Nil(t, out)
}

func TestDocsLLM_Handler(t *testing.T) {

	mockContent := "# I am a test article content"

	// Create mock fetcher that returns the mock docs
	mock := &MockDocsFetcher{
		FetchLLMDocsFunc: func(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{createDocResource(*articleURL, mockContent)},
			}, nil
		},
	}

	// Create the tool with disableResources set to true
	docsLLMTool := tool_docs_llm.NewDocsLLM(mock, false)

	// Execute the handler with empty request (should return default articles)
	request := &mcp.CallToolRequest{}
	in := tool_docs_llm.Input{
		ArticleURL: "https://some.firebolt.io/doc.md",
	}
	result, out, err := docsLLMTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	require.NotNil(t, out)
	assert.Len(t, out.Articles, 1)

	// Verify the content contains text content instead of embedded resources
	embContent, ok := out.Articles[0].(*mcp.EmbeddedResource)
	require.True(t, ok, "Expected EmbeddedResource when disableResources is false")
	assert.NotEmpty(t, embContent.Resource.Text)
	assert.Equal(t, embContent.Resource.URI, resources.LLMDocsURI(in.ArticleURL))
	assert.Equal(t, embContent.Resource.MIMEType, mimetype.Markdown)
	assert.Equal(t, mockContent, embContent.Resource.Text)
}

func TestDocsLLM_Handler_DisableResources(t *testing.T) {

	mockContent := "# I am a test article content"

	// Create mock fetcher that returns the mock docs
	mock := &MockDocsFetcher{
		FetchLLMDocsFunc: func(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error) {
			return &mcp.ReadResourceResult{
				Contents: []*mcp.ResourceContents{createDocResource(*articleURL, mockContent)},
			}, nil
		},
	}

	// Create the tool with disableResources set to true
	docsLLMTool := tool_docs_llm.NewDocsLLM(mock, true)

	// Execute the handler with empty request (should return default articles)
	request := &mcp.CallToolRequest{}
	in := tool_docs_llm.Input{
		ArticleURL: "https://some.firebolt.io/doc.md",
	}
	result, out, err := docsLLMTool.Handler()(t.Context(), request, in)

	// Assertions
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError)

	require.NotNil(t, out)
	assert.Len(t, out.Articles, 1)

	// Verify the content contains text content instead of embedded resources
	textContent, ok := out.Articles[0].(*mcp.TextContent)
	require.True(t, ok, "Expected TextContent when disableResources is true")
	assert.NotEmpty(t, textContent.Text)

	assert.Equal(t, mockContent, textContent.Text)
}

type MockDocsFetcher struct {
	FetchLLMDocsFunc func(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error)
}

func (m *MockDocsFetcher) FetchLLMDocsResources(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error) {
	return m.FetchLLMDocsFunc(ctx, articleURL)
}

// Helper to create a doc resource
func createDocResource(articleURL, content string) *mcp.ResourceContents {
	return &mcp.ResourceContents{
		URI:      "firebolt://llm_docs/" + articleURL,
		MIMEType: mimetype.Markdown,
		Text:     content,
	}
}
