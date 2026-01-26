package tool_docs

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

type Input struct {
}

type Output struct {
	Articles []mcp.Content `json:"articles"`
}

// DocsResourcesFetcher defines the interface for retrieving documentation resources.
// Implementations should provide methods to fetch Firebolt documentation.
type DocsResourcesFetcher interface {
	// FetchDocsResources retrieves documentation content for a specified article.
	// If article is empty, the implementation should determine an appropriate default behavior.
	FetchDocsResources(_ context.Context, article string) (*mcp.ReadResourceResult, error)
}

// Docs represents a tool for fetching and returning Firebolt documentation.
// It provides access to documentation articles that explain Firebolt concepts and functionality.
type Docs struct {
	docsFetcher      DocsResourcesFetcher // Fetches documentation resources
	disableResources bool                 // Return text content instead of embedded resources
	isCore           bool                 // Flag indicating whether the tool is for Firebolt Core or Firebolt SaaS
}

// Tool returns the mcp.Tool definition for the Docs tool.
// This defines how the tool is represented in the MCP system.
func (t *Docs) Tool() *mcp.Tool {
	return &mcp.Tool{
		Name:  "firebolt_docs_overview",
		Title: "Firebolt Documentation Overview",
		Description: "Returns Firebolt documentation overview. " +
			"Use this tool when need to get general information about Firebolt or need to connect to and use Firebolt. " +
			"Firebolt differs significantly from other databases, so it's important to gather some initial information before providing accurate answers. " +
			"To search for detailed information about using Firebolt, query syntax, object types etc use `firebolt_docs_search` tool",
	}
}

// Register adds the Docs tool to the provided MCP server instance.
func (t *Docs) Register(s *mcp.Server) {
	mcp.AddTool(s, t.Tool(), t.Handler())
}

func (t *Docs) Handler() mcp.ToolHandlerFor[Input, *Output] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, *Output, error) {

		var results []*mcp.ResourceContents // Collection of fetched documentation resources

		articleIDs := []string{}

		if t.isCore {
			articleIDs = append(articleIDs, resources.DocsCoreArticleOverview) // General Firebolt Core overview
		} else {
			articleIDs = append(articleIDs, resources.DocsArticleOverview) // General Firebolt overview
		}

		articleIDs = append(articleIDs, resources.DocsArticleProof) // Contains proof value for connect tool

		// Fetch each requested article
		for _, value := range articleIDs {
			// Fetch the article resources
			articleResources, err := t.docsFetcher.FetchDocsResources(ctx, value)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to discover resources: %w", err)
			}
			results = append(results, articleResources.Contents...)
		}

		out := itertools.Map(results, func(i *mcp.ResourceContents) mcp.Content {
			return tools.TextOrResourceContent(t.disableResources, i)
		})

		return &mcp.CallToolResult{}, &Output{Articles: out}, nil
	}
}

// NewDocs creates a new instance of the Docs tool with the provided documentation fetcher.
// It requires an implementation for fetching documentation articles.
func NewDocs(docsFetcher DocsResourcesFetcher, disableResources bool, isCore bool) *Docs {
	return &Docs{
		docsFetcher:      docsFetcher,
		disableResources: disableResources,
		isCore:           isCore,
	}
}
