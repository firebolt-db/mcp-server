package tool_docs_llm

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/bartventer/httpcache/store/memcache"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

type Input struct {
	ArticleURL string `json:"article_url,omitempty" jsonschema:"Firebolt Documentation article url to retrieve. If not passed, the full index of articles will be fetched"`
}

type Output struct {
	Articles []mcp.Content `json:"articles"`
}

// LLMDocsResourcesFetcher defines the interface for retrieving documentation resources targeted for LLMs.
// Implementations should provide methods to fetch Firebolt documentation.
type LLMDocsResourcesFetcher interface {
	// FetchLLMDocsResources retrieves documentation content (targeted for LLMs) for a specified article.
	// If the articleURL is empty, the implementation should return the documentation index.
	FetchLLMDocsResources(ctx context.Context, articleURL *string) (*mcp.ReadResourceResult, error)
}

type DocsLLM struct {
	disableResources bool // Return text content instead of embedded resources
	httpClient       *http.Client
	docsFetcher      LLMDocsResourcesFetcher
}

func (t *DocsLLM) Tool() *mcp.Tool {
	return &mcp.Tool{
		Name:  "firebolt_docs_llm",
		Title: "Firebolt Documentation Index and Retrieval for LLMs",
		Description: "Returns Firebolt documentation articles prepared to be used by LLMs. " +
			"Use this tool without parameters to list index of all Firebolt documentation articles. " +
			"Use this tools with a an article URL to retrieve specific documentation article.",
	}
}

// Register adds the Docs tool to the provided MCP server instance.
func (t *DocsLLM) Register(s *mcp.Server) {
	mcp.AddTool(s, t.Tool(), t.Handler())
}

func (t *DocsLLM) Handler() mcp.ToolHandlerFor[Input, *Output] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, *Output, error) {

		var urlToFetch *string

		if input.ArticleURL != "" {
			urlToFetch = &input.ArticleURL
		}

		cont, err := t.docsFetcher.FetchLLMDocsResources(ctx, urlToFetch)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to fetch LLM docs resource: %w", err)
		}

		out := itertools.Map(cont.Contents, func(i *mcp.ResourceContents) mcp.Content {
			return tools.TextOrResourceContent(t.disableResources, i)
		})
		return &mcp.CallToolResult{}, &Output{Articles: out}, nil
	}
}

// NewDocsLLM creates a new instance of the DocsLLM tool.
func NewDocsLLM(docsFetcher LLMDocsResourcesFetcher, disableResources bool) *DocsLLM {
	return &DocsLLM{
		docsFetcher:      docsFetcher,
		disableResources: disableResources,
		httpClient:       &http.Client{}, // In the future we may want to add a caching transport to this client
	}
}
