package databasemock

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
)

var _ database.Pool = &PoolMock{}

func NewPoolMock() *PoolMock {
	return &PoolMock{
		GetConnectionFunc: func(params database.PoolParams) (database.Connection, error) {
			return NewConnectionMock(), nil
		},
		CloseFunc:       func() {},
		IsClosed:        &atomic.Bool{},
		connections:     make(map[string]database.Connection),
		connectionCalls: make(map[string]int),
	}
}

type PoolMock struct {
	GetConnectionFunc func(params database.PoolParams) (database.Connection, error)
	CloseFunc         func()
	IsClosed          *atomic.Bool
	CloseCount        int

	mu              sync.RWMutex
	connections     map[string]database.Connection
	connectionCalls map[string]int
}

func (p *PoolMock) GetConnection(params database.PoolParams) (database.Connection, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.IsClosed.Load() {
		return nil, database.ErrPoolClosed
	}

	key := keyForPoolParams(params)
	p.connectionCalls[key]++

	return p.GetConnectionFunc(params)
}

func (p *PoolMock) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.CloseCount++
	p.IsClosed.Store(true)
	p.CloseFunc()
}

// WithGetConnectionFunc allows setting a custom GetConnection function for testing
func (p *PoolMock) WithGetConnectionFunc(fn func(params database.PoolParams) (database.Connection, error)) *PoolMock {
	p.GetConnectionFunc = fn
	return p
}

// WithCloseFunc allows setting a custom close function for testing
func (p *PoolMock) WithCloseFunc(fn func()) *PoolMock {
	p.CloseFunc = fn
	return p
}

// IsPoolClosed returns the status of the pool
func (p *PoolMock) IsPoolClosed() bool {
	return p.IsClosed.Load()
}

// GetConnectionCallCount returns the number of times GetConnection was called with the given parameters
func (p *PoolMock) GetConnectionCallCount(params database.PoolParams) int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	key := keyForPoolParams(params)
	return p.connectionCalls[key]
}

// RegisterConnection allows registering a specific connection for specific parameters
func (p *PoolMock) RegisterConnection(params database.PoolParams, conn database.Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	key := keyForPoolParams(params)
	p.connections[key] = conn

	originalFunc := p.GetConnectionFunc
	p.GetConnectionFunc = func(reqParams database.PoolParams) (database.Connection, error) {
		reqKey := keyForPoolParams(params)
		if conn, ok := p.connections[reqKey]; ok {
			return conn, nil
		}
		return originalFunc(reqParams)
	}
}

// keyForPoolParams generates a key by using the underlying string values
// for the optional parameters.
func keyForPoolParams(params database.PoolParams) string {
	db := ""
	if params.DatabaseName != nil {
		db = *params.DatabaseName
	}
	engine := ""
	if params.EngineName != nil {
		engine = *params.EngineName
	}
	return fmt.Sprintf("%s-%s-%s", params.AccountName, db, engine)
}
