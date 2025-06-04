package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/resources"
)

// DocsResourcesFetcher defines the interface for retrieving documentation resources.
// Implementations should provide methods to fetch Firebolt documentation.
type DocsResourcesFetcher interface {
	// FetchDocsResources retrieves documentation content for a specified article.
	// If article is empty, the implementation should determine an appropriate default behavior.
	FetchDocsResources(_ context.Context, article string) ([]mcp.ResourceContents, error)
}

// Docs represents a tool for fetching and returning Firebolt documentation.
// It provides access to documentation articles that explain Firebolt concepts and functionality.
type Docs struct {
	docsFetcher      DocsResourcesFetcher // Fetches documentation resources
	disableResources bool                 // Return text content instead of embedded resources
}

// NewDocs creates a new instance of the Docs tool with the provided documentation fetcher.
// It requires an implementation for fetching documentation articles.
func NewDocs(docsFetcher DocsResourcesFetcher, disableResources bool) *Docs {
	return &Docs{
		docsFetcher:      docsFetcher,
		disableResources: disableResources,
	}
}

// Tool returns the mcp.Tool definition for the Docs tool.
// This defines how the tool is represented in the MCP system, including its name, description,
// and parameters it accepts.
func (t *Docs) Tool() mcp.Tool {
	return mcp.NewTool(
		"firebolt_docs",
		mcp.WithDescription(
			"Returns Firebolt documentation articles. "+
				"Use this tool whenever you asked a question about Firebolt or need to connect to and use Firebolt. "+
				"Firebolt differs significantly from other databases, so it's important to gather some initial information before providing accurate answers. "+
				"Calling this tool without any parameters will return an overview document containing essential Firebolt fundamentals, "+
				"an index of detailed documentation articles, and a secret value expected by `firebolt_connect` tool that confirms you have read the documentation. "+
				"To retrieve specific articles, call this tool with their corresponding IDs using the `articles` parameter.",
		),
		mcp.WithArray(
			"articles",
			mcp.Title("Article IDs"),
			mcp.Description("Identifiers of the articles to fetch from Firebolt documentation"),
			mcp.MinItems(1),
			mcp.Items(map[string]any{
				"type": "string",
			}),
		),
	)
}

// Handler processes tool invocation requests and returns Firebolt documentation articles.
// If no specific articles are requested, it returns a set of default articles including an overview,
// a proof of documentation reading (needed for the connect tool), and a reference article.
// If specific articles are requested via the "articles" parameter, it fetches and returns those.
func (t *Docs) Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	var (
		articleIDs []any                  // List of article IDs to fetch
		results    []mcp.ResourceContents // Collection of fetched documentation resources
	)

	// Extract article IDs from request parameters if provided
	val, ok := request.GetArguments()["articles"]
	if ok && val != nil {
		articleIDs = val.([]any)
	}

	// Default articles to return if none specified
	if len(articleIDs) == 0 {
		articleIDs = append(
			articleIDs,
			resources.DocsArticleOverview,  // General Firebolt overview
			resources.DocsArticleProof,     // Contains proof value for connect tool
			resources.DocsArticleReference, // Reference documentation
		)
	}

	// Fetch each requested article
	for _, value := range articleIDs {
		// Ensure value is a string
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("invalid type for article ID: %T", value)
		}

		// Fetch the article resources
		articleResources, err := t.docsFetcher.FetchDocsResources(ctx, strValue)
		if err != nil {
			return nil, fmt.Errorf("failed to discover resources: %w", err)
		}
		results = append(results, articleResources...)
	}

	// Return the results as embedded resources
	return &mcp.CallToolResult{
		Result: mcp.Result{},
		Content: itertools.Map(results, func(i mcp.ResourceContents) mcp.Content {
			return textOrResourceContent(t.disableResources, i)
		}),
		IsError: false,
	}, nil
}
