package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/helpers/args"
)

// Query represents a tool for executing SQL queries against Firebolt databases.
// It handles connecting to the appropriate account, database, and engine before
// executing the query and returning results.
type Query struct {
	dbPool database.Pool
}

// NewQuery creates a new instance of the Query tool with the provided database connection pool.
// The connection pool is responsible for managing connections to different Firebolt
// account/database/engine combinations.
func NewQuery(dbPool database.Pool) *Query {
	return &Query{
		dbPool: dbPool,
	}
}

// Tool returns the mcp.Tool definition for the Query tool.
// This defines how the tool is represented in the MCP system, including:
// - The tool's name and description
// - Required and optional parameters with their descriptions
func (t *Query) Tool() mcp.Tool {
	return mcp.NewTool(
		"firebolt_query",
		mcp.WithDescription("Execute an SQL query against Firebolt."),
		mcp.WithString(
			"query",
			mcp.Required(),
			mcp.Title("Query"),
			mcp.Description("SQL query to execute"),
		),
		mcp.WithString(
			"account",
			mcp.Required(),
			mcp.Title("Account name"),
			mcp.Description("Name of the account to connect to"),
		),
		mcp.WithString(
			"database",
			mcp.Title("Database name"),
			mcp.Description(`
				Name of the database to send the query to.
				If not provided, no database will be specified.
				This still allows you to manage Firebolt organization and account metadata.
			`),
		),
		mcp.WithString(
			"engine",
			mcp.Title("Engine name"),
			mcp.Description(`
				Name of the engine to use for query execution.
				If not provided, the system engine will be used.
				Please note that the system engine can only be used for metadata queries and operations (DDL and DCL).
				Metadata queries are those that configure your Firebolt organization and account, or define the schema of your data.
				It will reject any queries that affect actual data stored in database.
			`),
		),
	)
}

// Handler processes tool invocation requests and executes SQL queries against Firebolt.
// It performs the following steps:
// 1. Extracts and validates required/optional parameters from the request
// 2. Acquires a database connection from the pool using the specified parameters
// 3. Executes the query against the database
// 4. Returns the query results as JSON
func (t *Query) Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// Extract the required and optional arguments from the request
	requireds, err1 := args.Strings(request.Params.Arguments, "query", "account")
	optionals, err2 := args.MaybeStrings(request.Params.Arguments, "database", "engine")
	if err := errors.Join(err1, err2); err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	// Acquire a connection to the database using the provided parameters
	conn, err := t.dbPool.GetConnection(database.PoolParams{
		AccountName:  requireds[1],
		DatabaseName: optionals[0],
		EngineName:   optionals[1],
	})
	if err != nil {
		return nil, fmt.Errorf("failed to obtain database connection: %w", err)
	}

	// Execute the query using the connection
	result, err := conn.Query(ctx, requireds[0])
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	// Convert the query result to JSON format
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query result: %w", err)
	}

	// Create a new tool result with the JSON data and return it
	return mcp.NewToolResultText(string(resultJSON)), nil
}
