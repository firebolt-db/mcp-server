package resources_test

import (
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
	"github.com/firebolt-db/mcp-server/pkg/resources"
)

const (
	llmsIndexURL = "https://docs.firebolt.io/llms.txt"
)

func TestLLMDocs_ResourceTemplate(t *testing.T) {
	llmDocs := resources.NewLLMDocs()
	template := llmDocs.ResourceTemplate()

	assert.NotEmpty(t, template.Name)
	assert.NotEmpty(t, template.URITemplate)
	assert.NotEmpty(t, template.MIMEType)
	assert.NotEmpty(t, template.Description)
}

func TestLLMDocs_Handler(t *testing.T) {
	t.Parallel()

	llmDocs := resources.NewLLMDocs()

	t.Run("fetch docs index", func(t *testing.T) {
		t.Parallel()

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"articleURL": llmsIndexURL,
				},
			},
		}

		result, err := llmDocs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.LLMDocsURI(llmsIndexURL), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Contains(t, resource.Text, "# Firebolt Documentation")
	})

	t.Run("fetch single article", func(t *testing.T) {
		t.Parallel()

		articleURL := "https://docs.firebolt.io/reference-sql/functions-reference/string/ltrim.md"

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"articleURL": articleURL,
				},
			},
		}

		result, err := llmDocs.Handler(t.Context(), request)
		assert.NoError(t, err)
		require.Len(t, result.Contents, 1)
		resource := result.Contents[0]
		assert.Equal(t, resources.LLMDocsURI(articleURL), resource.URI)
		assert.Equal(t, mimetype.Markdown, resource.MIMEType)
		assert.Contains(t, resource.Text, "> Reference material for LTRIM function")
	})

	t.Run("fetch non-existent article", func(t *testing.T) {
		t.Parallel()

		articleURL := "https://docs.firebolt.io/non-existent-article.md"

		request := &mcp.ReadResourceRequest{
			Params: &mcp.ReadResourceParams{
				Meta: map[string]any{
					"articleURL": articleURL,
				},
			},
		}

		_, err := llmDocs.Handler(t.Context(), request)
		assert.Error(t, err)
	})
}
