package tool_query

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/tools"
)

type Input struct {
	Query    string `json:"query" jsonschema:"SQL query to execute"`
	Account  string `json:"account" jsonschema:"Name of the Firebolt account to connect to"`
	Database string `json:"database,omitempty" jsonschema:"Name of the database to send the query to. If not provided, no database will be specified. This still allows you to manage Firebolt organization and account metadata."`
	Engine   string `json:"engine,omitempty" jsonschema:"Name of the engine to use for query execution. If not provided, the system engine will be used. Please note that the system engine can only be used for metadata queries and operations (DDL and DCL). Metadata queries are those that configure your Firebolt organization and account, or define the schema of your data. It will reject any queries that affect actual data stored in database."`
}

type Output struct {
	Result []mcp.Content `json:"result"`
}

// Query represents a tool for executing SQL queries against Firebolt databases.
// It handles connecting to the appropriate account, database, and engine before
// executing the query and returning results.
type Query struct {
	dbPool   database.Pool
	readOnly bool
}

// Option configures the Query tool.
type Option func(*Query)

// WithReadOnly restricts the tool to a single SELECT or WITH statement and
// declares it read-only in its MCP annotations, so clients may cache its results.
func WithReadOnly() Option {
	return func(q *Query) {
		q.readOnly = true
	}
}

// NewQuery creates a new instance of the Query tool with the provided database connection pool.
// The connection pool is responsible for managing connections to different Firebolt
// account/database/engine combinations.
func NewQuery(dbPool database.Pool, opts ...Option) *Query {
	q := &Query{
		dbPool: dbPool,
	}
	for _, opt := range opts {
		opt(q)
	}
	return q
}

// Tool returns the mcp.Tool definition for the Query tool.
// This defines how the tool is represented in the MCP system, including:
// - The tool's name and description
// - Required and optional parameters with their descriptions
func (t *Query) Tool() *mcp.Tool {
	tool := &mcp.Tool{
		Name:        "firebolt_query",
		Title:       "Firebolt Query",
		Description: "Execute an SQL query against Firebolt.",
	}
	if t.readOnly {
		tool.Description += " Read-only: only a single SELECT or WITH statement is accepted."
		tool.Annotations = &mcp.ToolAnnotations{
			ReadOnlyHint:    true,
			DestructiveHint: tools.Ptr(false),
			IdempotentHint:  true,
			OpenWorldHint:   tools.Ptr(false),
		}
	}
	return tool
}

// Register adds the Query tool to the provided MCP server instance.
func (t *Query) Register(s *mcp.Server) {
	mcp.AddTool(s, t.Tool(), t.Handler())
}

// Handler processes tool invocation requests and executes SQL queries against Firebolt.
// It performs the following steps:
// 1. Extracts and validates required/optional parameters from the request
// 2. Acquires a database connection from the pool using the specified parameters
// 3. Executes the query against the database
// 4. Returns the query results as JSON
func (t *Query) Handler() mcp.ToolHandlerFor[Input, *Output] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, *Output, error) {
		if input.Query == "" {
			return nil, nil, fmt.Errorf("bad request: query parameter is required")
		}
		if input.Account == "" {
			return nil, nil, fmt.Errorf("bad request: account parameter is required")
		}
		if t.readOnly {
			if err := validateReadOnly(ctx, input.Query); err != nil {
				return nil, nil, err
			}
		}

		var (
			dbName *string
			engine *string
		)

		if input.Database != "" {
			dbName = &input.Database
		}
		if input.Engine != "" {
			engine = &input.Engine
		}

		// Acquire a connection to the database using the provided parameters
		conn, err := t.dbPool.GetConnection(database.PoolParams{
			AccountName:  input.Account,
			DatabaseName: dbName,
			EngineName:   engine,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to obtain database connection: %w", err)
		}

		// Execute the query using the connection
		result, err := conn.Query(ctx, input.Query)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to query database: %w", err)
		}

		// Convert the query result to JSON format
		resultJSON, err := json.Marshal(result)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal query result: %w", err)
		}

		out := &Output{
			Result: []mcp.Content{
				&mcp.TextContent{
					Text: string(resultJSON),
				},
			},
		}

		return &mcp.CallToolResult{}, out, nil
	}
}
