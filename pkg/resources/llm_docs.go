package resources

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/helpers/args"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
)

const (
	// url where the documentation index for the LLMs is located
	llmsIndexURL = "https://docs.firebolt.io/llms.txt"
)

func LLMDocsURI(url string) string {
	return fmt.Sprintf("firebolt://llm_docs/%s", url)
}

// LLMDocs is a resource template handler for serving Firebolt documentation articles for LLMs.
type LLMDocs struct {
	httpClient *http.Client
}

// NewLLMDocs creates and returns a new instance of the DocsLLM.
func NewLLMDocs() *LLMDocs {
	return &LLMDocs{
		httpClient: &http.Client{}, // In the future we may want to add a caching transport to this client
	}
}

// ResourceTemplate defines the template for LLM documentation resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *LLMDocs) ResourceTemplate() *mcp.ResourceTemplate {
	return &mcp.ResourceTemplate{
		URITemplate: LLMDocsURI("{url}"),
		Name:        "Documentation article for LLM",
		MIMEType:    mimetype.Markdown,
		Description: "An article that offers insights into a particular topic related to Firebolt, specifically compiled for LLMs.",
		Annotations: &mcp.Annotations{
			Audience: []mcp.Role{"user", "assistant"},
			Priority: 0.5,
		},
	}
}

// Handler processes resource requests for LLM documentation articles.
// It extracts the article URL parameter and fetches the appropriate documentation content.
func (r *LLMDocs) Handler(ctx context.Context, request *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	// Extract the article URL from the request parameters
	value, err := args.String(request.GetParams().GetMeta(), "articleURL")
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	return r.FetchLLMDocsResources(ctx, &value)
}

// FetchLLMDocsResources retrieves the content for a specified documentation article.
// If the provided articleURL is nil, it returns the documentation index.
func (r *LLMDocs) FetchLLMDocsResources(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error) {
	if articleURL == nil {
		return r.fetchArticle(ctx, llmsIndexURL)
	} else {
		return r.fetchArticle(ctx, *articleURL)
	}
}

func (r *LLMDocs) fetchArticle(ctx context.Context, articleURL string) (*mcp.ReadResourceResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, articleURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make http request: %w", err)
	}
	defer resp.Body.Close()

	var responseText string
	responseData, err := io.ReadAll(resp.Body)
	if err == nil {
		responseText = string(responseData)
	}

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected response status code returned from Firebolt Docs: code - %d, responseText - %s",
			resp.StatusCode, responseText,
		)
	}

	result := &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{
				URI:      LLMDocsURI(articleURL),
				MIMEType: mimetype.Markdown,
				Text:     responseText,
			},
		},
	}

	return result, nil
}
