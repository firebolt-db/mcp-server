package databasemock

import (
	"context"
	"sync/atomic"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
)

var _ database.ConnectionCloser = &ConnectionMock{}

func NewConnectionMock() *ConnectionMock {
	return &ConnectionMock{
		QueryFunc: func(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
			return []map[string]any{}, nil
		},
		CloseFunc: func() {},
		IsClosed:  &atomic.Bool{},
	}
}

type ConnectionMock struct {
	QueryFunc  func(ctx context.Context, sql string, args ...any) ([]map[string]any, error)
	CloseFunc  func()
	CloseCount int
	IsClosed   *atomic.Bool
}

func (c *ConnectionMock) Query(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {
	if c.IsClosed.Load() {
		return nil, database.ErrConnectionClosed
	}
	return c.QueryFunc(ctx, sql, args...)
}

func (c *ConnectionMock) Close() {
	c.CloseCount++
	c.IsClosed.Store(true)
	c.CloseFunc()
}

// WithQueryFunc allows setting a custom query function for testing
func (c *ConnectionMock) WithQueryFunc(fn func(ctx context.Context, sql string, args ...any) ([]map[string]any, error)) *ConnectionMock {
	c.QueryFunc = fn
	return c
}

// WithCloseFunc allows setting a custom close function for testing
func (c *ConnectionMock) WithCloseFunc(fn func()) *ConnectionMock {
	c.CloseFunc = fn
	return c
}

// IsConnectionClosed returns the status of the connection
func (c *ConnectionMock) IsConnectionClosed() bool {
	return c.IsClosed.Load()
}
