package tool_connect

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/sync/errgroup"

	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

type Input struct {
	DocsProof string `json:"docs_proof" jsonschema:"This parameter is used to confirm that the essential documentation has been reviewed before connecting to Firebolt. The correct value will be returned by the 'firebolt_docs' tool when it is called without any parameters."`
}

type Output struct {
	Results []mcp.Content `json:"results"`
}

// AccountResourcesFetcher defines the interface for retrieving account resources.
// Implementations should provide methods to fetch Firebolt account information.
type AccountResourcesFetcher interface {
	// FetchAccountResources retrieves account information for a specified account name.
	// If accountName is empty, all accessible accounts should be returned.
	FetchAccountResources(ctx context.Context, accountName string) (*mcp.ReadResourceResult, error)
}

// DatabaseResourcesFetcher defines the interface for retrieving database resources.
// Implementations should provide methods to fetch Firebolt database information within an account.
type DatabaseResourcesFetcher interface {
	// FetchDatabaseResources retrieves database information for a specified account and database name.
	// If databaseName is empty, all databases in the account should be returned.
	FetchDatabaseResources(ctx context.Context, accountName, databaseName string) (*mcp.ReadResourceResult, error)
}

// EngineResourcesFetcher defines the interface for retrieving engine resources.
// Implementations should provide methods to fetch Firebolt engine information within an account.
type EngineResourcesFetcher interface {
	// FetchEngineResources retrieves engine information for a specified account and engine name.
	// If engineName is empty, all engines in the account should be returned.
	FetchEngineResources(ctx context.Context, accountName, engineName string) (*mcp.ReadResourceResult, error)
}

// Connect represents a tool for fetching and returning Firebolt resource information.
// It provides hierarchical access to accounts, databases, and engines in the Firebolt system.
type Connect struct {
	accountsFetcher  AccountResourcesFetcher  // Fetches account resources
	databasesFetcher DatabaseResourcesFetcher // Fetches database resources
	enginesFetcher   EngineResourcesFetcher   // Fetches engine resources
	docsProof        string                   // Shared with the docs resources
	disableResources bool                     // Return text content instead of embedded resources
}

// NewConnect creates a new instance of the Connect tool with the provided resource fetchers.
// It requires implementations for fetching accounts, databases, and engines.
func NewConnect(
	accountsFetcher AccountResourcesFetcher,
	databasesFetcher DatabaseResourcesFetcher,
	enginesFetcher EngineResourcesFetcher,
	docsProof string,
	disableResources bool,
) *Connect {
	return &Connect{
		accountsFetcher:  accountsFetcher,
		databasesFetcher: databasesFetcher,
		enginesFetcher:   enginesFetcher,
		docsProof:        docsProof,
		disableResources: disableResources,
	}
}

// Tool returns the mcp.Tool definition for the Connect tool.
// This defines how the tool is represented in the MCP system.
func (t *Connect) Tool() *mcp.Tool {
	return &mcp.Tool{
		Name:  "firebolt_connect",
		Title: "Firebolt Connect",
		Description: "Returns a list of Firebolt accounts, databases, and engines you have access to. " +
			"This information is required before using the `firebolt_query` tool.",
	}
}

// Register adds the Connect tool to the provided MCP server instance.
func (t *Connect) Register(s *mcp.Server) {
	mcp.AddTool(s, t.Tool(), t.Handler())
}

// Handler processes tool invocation requests and returns a comprehensive view of
// Firebolt resources. It fetches accounts and then concurrently retrieves the
// databases and engines for each account.
func (t *Connect) Handler() mcp.ToolHandlerFor[Input, *Output] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, *Output, error) {
		if input.DocsProof != t.docsProof {
			return nil, nil, fmt.Errorf("invalid documentation proof, " +
				"you need to call `firebolt_docs` tool first and extract value from this parameter from the response")
		}

		// Fetch all accounts
		accounts, err := t.accountsFetcher.FetchAccountResources(ctx, "")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to discover resources: %w", err)
		}

		if len(accounts.Contents) == 0 {
			return nil, nil, errors.New("your identity does not have access to any Firebolt accounts")
		}

		// Results collection
		var results []*mcp.ResourceContents
		done := make(chan struct{}, 1)
		resourcesCh := make(chan *mcp.ResourceContents, len(accounts.Contents))

		// Start a goroutine to collect results from the channel
		go func() {
			for resource := range resourcesCh {
				results = append(results, resource)
			}
			done <- struct{}{}
		}()

		// For every account, concurrently fetch databases and engines
		group, groupCtx := errgroup.WithContext(ctx)
		for _, account := range accounts.Contents {
			account := account // Capture variable for goroutine
			group.Go(func() error {

				// Add the account resource to results
				resourcesCh <- account

				// Extract the account information from the resource

				var ar accountResource
				err := json.Unmarshal([]byte(account.Text), &ar)
				if err != nil {
					return fmt.Errorf("failed to unmarshal account resource: %w", err)
				}

				// Fetch all databases for the account
				databases, err := t.databasesFetcher.FetchDatabaseResources(groupCtx, ar.Name, "")
				if err != nil {
					return fmt.Errorf("failed to discover database resources: %w", err)
				}
				for _, db := range databases.Contents {
					resourcesCh <- db
				}

				// Fetch all engines for the account
				engines, err := t.enginesFetcher.FetchEngineResources(groupCtx, ar.Name, "")
				if err != nil {
					return fmt.Errorf("failed to discover engine resources: %w", err)
				}
				for _, engine := range engines.Contents {
					resourcesCh <- engine
				}

				return nil
			})
		}

		// Wait for all fetches to complete
		err = group.Wait()
		close(resourcesCh)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to fetch resources: %w", err)
		}

		// Wait for all results to be collected
		<-done

		out := itertools.Map(results, func(i *mcp.ResourceContents) mcp.Content {
			return tools.TextOrResourceContent(t.disableResources, i)
		})

		return &mcp.CallToolResult{}, &Output{Results: out}, nil
	}
}

// accountResource represents the minimal structure needed to extract
// the account name from an account resource.
type accountResource struct {
	Name string `json:"name"`
}
