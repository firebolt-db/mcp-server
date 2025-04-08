package discoverymock

import (
	"context"

	"github.com/firebolt-db/mcp-server/pkg/clients/discovery"
)

var _ discovery.Client = &ClientMock{}

// ClientMock is a mock implementation of the discovery.Client interface
type ClientMock struct {
	ListAccountsFunc  func(ctx context.Context) ([]discovery.Account, error)
	ListAccountsCount int
}

// NewClientMock creates a new instance of ClientMock with default implementations
func NewClientMock() *ClientMock {
	return &ClientMock{
		ListAccountsFunc: func(ctx context.Context) ([]discovery.Account, error) {
			return []discovery.Account{}, nil
		},
	}
}

// ListAccounts implements the discovery.Client interface
func (c *ClientMock) ListAccounts(ctx context.Context) ([]discovery.Account, error) {
	c.ListAccountsCount++
	return c.ListAccountsFunc(ctx)
}

// WithListAccountsFunc allows setting a custom ListAccounts function for testing
func (c *ClientMock) WithListAccountsFunc(fn func(ctx context.Context) ([]discovery.Account, error)) *ClientMock {
	c.ListAccountsFunc = fn
	return c
}
