package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/firebolt-db/mcp-server/pkg/clients/discovery"
	"github.com/firebolt-db/mcp-server/pkg/helpers/args"
	"github.com/firebolt-db/mcp-server/pkg/helpers/itertools"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
)

// AccountURI creates a formatted Firebolt account URI for a given account name.
func AccountURI(name string) string {
	return fmt.Sprintf("firebolt://accounts/%s", name)
}

// Accounts is a resource template handler for serving Firebolt account information.
// It interacts with the discovery client to retrieve account data from the Firebolt service.
type Accounts struct {
	discoveryClient discovery.Client
}

// NewAccounts creates and returns a new instance of the Accounts resource handler
// with the provided discovery client.
func NewAccounts(discoveryClient discovery.Client) *Accounts {
	return &Accounts{
		discoveryClient: discoveryClient,
	}
}

// ResourceTemplate defines the template for account resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *Accounts) ResourceTemplate() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		AccountURI("{account}"),
		"Account",
		mcp.WithTemplateMIMEType(mimetype.JSON),
		mcp.WithTemplateAnnotations([]mcp.Role{mcp.RoleUser, mcp.RoleAssistant}, 0.9),
		mcp.WithTemplateDescription("Brief information about the account in the Firebolt organization."),
	)
}

// Handler processes resource requests for account information.
// It extracts the account parameter and fetches the appropriate account data.
func (r *Accounts) Handler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {

	accountName, err := args.String(request.Params.Arguments, "account")
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	return r.FetchAccountResources(ctx, accountName)
}

// FetchAccountResources retrieves account information from the Firebolt service.
// If a specific account name is specified, it filters for that account; otherwise, it returns all accounts.
func (r *Accounts) FetchAccountResources(ctx context.Context, accountName string) ([]mcp.ResourceContents, error) {

	// Fetch the list of accounts from the discovery client
	accounts, err := r.discoveryClient.ListAccounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to discover accounts: %w", err)
	}

	// Filter the accounts by name if needed
	var filteredAccounts []discovery.Account
	for _, account := range accounts {
		if accountName == "" || accountName == account.Name {
			filteredAccounts = append(filteredAccounts, account)
		}
	}

	return itertools.MapWithFailure(filteredAccounts, func(i discovery.Account) (mcp.ResourceContents, error) {

		data, err := json.Marshal(i)
		if err != nil {
			return mcp.TextResourceContents{}, fmt.Errorf("failed to marshal row data to JSON: %w", err)
		}

		return mcp.TextResourceContents{
			URI:      AccountURI(i.Name),
			MIMEType: mimetype.JSON,
			Text:     string(data),
		}, nil
	})
}
