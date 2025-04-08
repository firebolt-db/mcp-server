package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/helpers/args"
	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
)

// DatabaseURI creates a formatted Firebolt database URI for a given account and database name.
func DatabaseURI(account, database string) string {
	return fmt.Sprintf("firebolt://accounts/%s/databases/%s", account, database)
}

// Databases is a resource template handler for serving Firebolt database information.
type Databases struct {
	dbPool database.Pool
}

// NewDatabases creates and returns a new instance of the Databases resource handler.
func NewDatabases(dbPool database.Pool) *Databases {
	return &Databases{
		dbPool: dbPool,
	}
}

// ResourceTemplate defines the template for database resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *Databases) ResourceTemplate() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		DatabaseURI("{account}", "{database}"),
		"Database",
		mcp.WithTemplateMIMEType(mimetype.JSON),
		mcp.WithTemplateAnnotations([]mcp.Role{mcp.RoleUser, mcp.RoleAssistant}, 0.8),
		mcp.WithTemplateDescription("Brief information about the database in the Firebolt account."),
	)
}

// Handler processes resource requests for database information.
// It extracts account and database parameters and fetches the appropriate database data.
func (r *Databases) Handler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {

	values, err := args.Strings(request.Params.Arguments, "account", "database")
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	return r.FetchDatabaseResources(ctx, values[0], values[1])
}

// FetchDatabaseResources retrieves database information from the Firebolt service.
// If a specific database is specified, it filters for that database; otherwise, it returns all databases.
func (r *Databases) FetchDatabaseResources(ctx context.Context, account, dbName string) ([]mcp.ResourceContents, error) {

	// Acquire a connection to the database
	conn, err := r.dbPool.GetConnection(database.PoolParams{
		AccountName: account,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection: %w", err)
	}

	// Prepare the SQL query
	var (
		sql  = "SELECT database_name, description FROM information_schema.databases"
		args []any
	)
	if dbName != "" {
		sql += " WHERE database_name = ?"
		args = append(args, dbName)
	}

	// Query database
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	// Convert rows to resources
	return itertools.MapWithFailure(rows, func(i map[string]any) (mcp.ResourceContents, error) {

		i["account_name"] = account
		data, err := json.Marshal(i)
		if err != nil {
			return mcp.TextResourceContents{}, fmt.Errorf("failed to marshal row data to JSON: %w", err)
		}

		return mcp.TextResourceContents{
			URI:      DatabaseURI(account, i["database_name"].(string)),
			MIMEType: mimetype.JSON,
			Text:     string(data),
		}, nil
	})
}
