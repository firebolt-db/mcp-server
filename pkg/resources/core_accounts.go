package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
)

// CoreAccounts is a resource template handler for serving Firebolt Core account information.
// It interacts with the database pool to retrieve account data from the Firebolt Core.
type CoreAccounts struct {
	dbPool database.Pool
}

// NewCoreAccounts creates and returns a new instance of the CoreAccounts resource handler
// with the provided dbPool.
func NewCoreAccounts(dbPool database.Pool) *CoreAccounts {
	return &CoreAccounts{
		dbPool: dbPool,
	}
}

// ResourceTemplate defines the template for CoreAccount resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *CoreAccounts) ResourceTemplate() *mcp.ResourceTemplate {
	return &mcp.ResourceTemplate{
		URITemplate: AccountURI("{account}"),
		Name:        "Account",
		MIMEType:    mimetype.JSON,
		Description: "Brief information about the account in Firebolt Core.",
		Annotations: &mcp.Annotations{
			Audience: []mcp.Role{"user", "assistant"},
			Priority: 0.9,
		},
	}
}

// Handler processes resource requests for account information.
// It extracts the account parameter and fetches the appropriate account data.
func (r *CoreAccounts) Handler(ctx context.Context, request *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	return r.FetchAccountResources(ctx)
}

// FetchAccountResources retrieves account information from the Firebolt service.
// If a specific account name is specified, it filters for that account; otherwise, it returns all accounts.
func (r *CoreAccounts) FetchAccountResources(ctx context.Context) (*mcp.ReadResourceResult, error) {

	// Acquire a connection to the database
	conn, err := r.dbPool.GetConnection(database.PoolParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to acquire database connection: %w", err)
	}

	sql := "SELECT account_id, account_name FROM information_schema.accounts;"

	// run the query
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
			URI:      AccountURI(i["account_name"].(string)),
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
