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

// EngineURI creates a formatted Firebolt engine URI for a given account and engine name.
func EngineURI(account, engine string) string {
	return fmt.Sprintf("firebolt://accounts/%s/engines/%s", account, engine)
}

// Engines is a resource template handler for serving Firebolt engine information.
type Engines struct {
	dbPool database.Pool
}

// NewEngines creates and returns a new instance of the Engines resource handler.
func NewEngines(dbPool database.Pool) *Engines {
	return &Engines{
		dbPool: dbPool,
	}
}

// ResourceTemplate defines the template for engine resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *Engines) ResourceTemplate() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		EngineURI("{account}", "{engine}"),
		"Engine",
		mcp.WithTemplateMIMEType(mimetype.JSON),
		mcp.WithTemplateAnnotations([]mcp.Role{mcp.RoleUser, mcp.RoleAssistant}, 0.8),
		mcp.WithTemplateDescription("Brief information about the engine in the Firebolt account."),
	)
}

// Handler processes resource requests for engine information.
// It extracts account and engine parameters and fetches the appropriate engine data.
func (r *Engines) Handler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {

	params, err := args.Strings(request.Params.Arguments, "account", "engine")
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	return r.FetchEngineResources(ctx, params[0], params[1])
}

// FetchEngineResources retrieves engine information from the database.
// If a specific engine is specified, it filters for that engine; otherwise, it returns all engines.
func (r *Engines) FetchEngineResources(ctx context.Context, account, engine string) ([]mcp.ResourceContents, error) {

	// Acquire a connection to the database
	conn, err := r.dbPool.GetConnection(database.PoolParams{
		AccountName: account,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection: %w", err)
	}

	// Prepare the SQL query
	var (
		sql  = "SELECT engine_name, description, status, version, type, family, nodes, clusters, auto_start FROM information_schema.engines"
		args []any
	)
	if engine != "" {
		sql += " WHERE engine_name = ?"
		args = append(args, engine)
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
			URI:      EngineURI(account, i["engine_name"].(string)),
			MIMEType: mimetype.JSON,
			Text:     string(data),
		}, nil
	})
}
