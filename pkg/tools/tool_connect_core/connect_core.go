package tool_connect_core

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

type Input struct {
	DocsProof string `json:"docs_proof,omitempty" jsonschema:"This parameter is used to confirm that the essential documentation has been reviewed before connecting to Firebolt. The correct value will be returned by the 'firebolt_docs_overview' tool when it is called without any parameters."`
}

type Output struct {
	Results []mcp.Content `json:"results"`
}

// CoreDatabaseResourcesFetcher defines the interface for retrieving database resources in Firebolt Core.
// Implementations should provide methods to fetch Firebolt database information in Firebolt Core.
type CoreDatabaseResourcesFetcher interface {
	// FetchCoreDatabaseResources retrieves database information in a Firebolt Core instance.
	FetchCoreDatabaseResources(ctx context.Context) (*mcp.ReadResourceResult, error)
}

// ConnectCore represents a tool for fetching and returning Firebolt resource information.
// It provides hierarchical databases in the Firebolt Core system.
type ConnectCore struct {
	databasesFetcher CoreDatabaseResourcesFetcher // Fetches database resources
	docsProof        *string                      // Shared with the docs resources
	disableResources bool                         // Return text content instead of embedded resources
}

// NewConnectCore creates a new instance of the Connect tool with the provided resource fetchers.
// It requires implementation for fetching databases.
func NewConnectCore(
	databasesFetcher CoreDatabaseResourcesFetcher,
	docsProof *string,
	disableResources bool,
) *ConnectCore {
	return &ConnectCore{
		databasesFetcher: databasesFetcher,
		docsProof:        docsProof,
		disableResources: disableResources,
	}
}

// Tool returns the mcp.Tool definition for the ConnectCore tool.
// This defines how the tool is represented in the MCP system.
func (t *ConnectCore) Tool() *mcp.Tool {
	return &mcp.Tool{
		Name:  "firebolt_connect_core",
		Title: "Firebolt Core Connect",
		Description: "Returns a list of Firebolt databases you have access to. " +
			"This information is required before using the `firebolt_query` tool.",
	}
}

// Register adds the Connect tool to the provided MCP server instance.
func (t *ConnectCore) Register(s *mcp.Server) {
	mcp.AddTool(s, t.Tool(), t.Handler())
}

// Handler processes tool invocation requests and returns a comprehensive view of
// Firebolt resources. It fetches accounts and then concurrently retrieves the
// databases and engines for each account.
func (t *ConnectCore) Handler() mcp.ToolHandlerFor[Input, *Output] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, *Output, error) {
		if t.docsProof != nil && input.DocsProof != *t.docsProof {
			return nil, nil, fmt.Errorf("invalid documentation proof, " +
				"you need to call `firebolt_docs_overview` tool first and extract value from this parameter from the response")
		}

		// Fetch all databases for the account
		databases, err := t.databasesFetcher.FetchCoreDatabaseResources(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to discover database resources: %w", err)
		}

		out := itertools.Map(databases.Contents, func(i *mcp.ResourceContents) mcp.Content {
			return tools.TextOrResourceContent(t.disableResources, i)
		})

		return &mcp.CallToolResult{}, &Output{Results: out}, nil
	}
}
