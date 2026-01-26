package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/helpers/args"
	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
)

// EngineURI creates a formatted Firebolt engine URI for a given account and engine name.
func EngineURI(account, engine string) string {
	return fmt.Sprintf("firebolt://accounts/%s/engines/%s", account, engine)
}

// CoreEngineURI creates a formatted Firebolt Core engine URI for a given engine name.
func CoreEngineURI(engine string) string {
	return fmt.Sprintf("firebolt://engines/%s", engine)
}

// Engines is a resource template handler for serving Firebolt engine information.
type Engines struct {
	dbPool database.Pool
	isCore bool
}

// NewEngines creates and returns a new instance of the Engines resource handler.
func NewEngines(dbPool database.Pool, opts ...resourceOption) *Engines {
	optSet := resourceOptionSet{}
	for _, opt := range opts {
		opt(&optSet)
	}

	return &Engines{
		dbPool: dbPool,
		isCore: optSet.isCore,
	}
}

// ResourceTemplate defines the template for engine resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *Engines) ResourceTemplate() *mcp.ResourceTemplate {
	uriTemplate := EngineURI("{account}", "{engine}")
	if r.isCore {
		uriTemplate = CoreEngineURI("{engine}")
	}

	return &mcp.ResourceTemplate{
		URITemplate: uriTemplate,
		Name:        "Engine",
		MIMEType:    mimetype.JSON,
		Description: "Brief information about the engine in the Firebolt account.",
		Annotations: &mcp.Annotations{
			Audience: []mcp.Role{"user", "assistant"},
			Priority: 0.8,
		},
	}
}

// Handler processes resource requests for engine information.
// It extracts account and engine parameters and fetches the appropriate engine data.
func (r *Engines) Handler(ctx context.Context, request *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {

	var accountName, engineName string

	if !r.isCore {
		values, err := args.Strings(request.GetParams().GetMeta(), "account", "engine")
		if err != nil {
			return nil, fmt.Errorf("bad request: %w", err)
		}

		accountName, engineName = values[0], values[1]
	}

	return r.FetchEngineResources(ctx, accountName, engineName)
}

// FetchEngineResources retrieves engine information from the database.
// If a specific engine is specified, it filters for that engine; otherwise, it returns all engines.
func (r *Engines) FetchEngineResources(ctx context.Context, account, engine string) (*mcp.ReadResourceResult, error) {
	if r.isCore {
		return r.fetchCoreEngineResources(ctx)
	}

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
	out, err := itertools.MapWithFailure(rows, func(i map[string]any) (*mcp.ResourceContents, error) {

		i["account_name"] = account
		data, err := json.Marshal(i)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal row data to JSON: %w", err)
		}

		return &mcp.ResourceContents{
			URI:      EngineURI(account, i["engine_name"].(string)),
			MIMEType: mimetype.JSON,
			Text:     string(data),
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: out,
	}, nil
}

// fetchCoreEngineResources retrieves engine information in Firebolt Core.
func (r *Engines) fetchCoreEngineResources(ctx context.Context) (*mcp.ReadResourceResult, error) {

	// Acquire a connection to the database
	conn, err := r.dbPool.GetConnection(database.PoolParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection: %w", err)
	}

	// Prepare the SQL query
	sql := "SELECT engine_name, description, status, version, type, family, nodes, clusters, auto_start FROM information_schema.engines;"

	// Query database
	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	// Convert rows to resources
	out, err := itertools.MapWithFailure(rows, func(i map[string]any) (*mcp.ResourceContents, error) {

		data, err := json.Marshal(i)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal row data to JSON: %w", err)
		}

		return &mcp.ResourceContents{
			URI:      CoreEngineURI(i["engine_name"].(string)),
			MIMEType: mimetype.JSON,
			Text:     string(data),
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: out,
	}, nil
}
