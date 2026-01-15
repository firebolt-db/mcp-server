package database

import (
	"fmt"
	"log/slog"
	"sync"
)

// NewCorePool creates a new connection pool to Firebolt Core with the provided logger.
// It returns the pool implementation and a function to close all connections.
func NewCorePool(logger *slog.Logger, url string) (Pool, func()) {
	return NewCorePoolWithConnectionFactory(logger, url, NewConnection)
}

// NewCorePoolWithConnectionFactory creates a new connection pool to Firebolt Core with a custom connection factory.
// It allows you to provide a custom function to create connections which can be useful for testing.
func NewCorePoolWithConnectionFactory(logger *slog.Logger, url string, newConnectionFunc NewConnectionFunc) (Pool, func()) {
	pool := &corePoolImpl{
		url:               url,
		logger:            logger,
		connections:       make(map[string]Connection),
		newConnectionFunc: newConnectionFunc,
	}
	return pool, pool.Close
}

type corePoolImpl struct {
	sync.Mutex
	url               string
	isClosed          bool
	logger            *slog.Logger
	closers           []func()
	connections       map[string]Connection
	newConnectionFunc NewConnectionFunc
}

func (p *corePoolImpl) GetConnection(params PoolParams) (Connection, error) {
	p.Lock()
	defer p.Unlock()

	connectionParams := CoreConnectionParams{
		URL:          p.url,
		DatabaseName: params.DatabaseName,
	}
	hash := connectionParams.Hash()

	// First, try to get an existing connection
	if p.isClosed {
		return nil, ErrPoolClosed
	}
	if conn, ok := p.connections[hash]; ok {
		return conn, nil
	}

	// Create a new connection if one doesn't exist
	conn, closer, err := p.newConnectionFunc(p.logger, connectionParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	// Store the new connection in the pool
	p.connections[hash] = conn
	p.closers = append(p.closers, closer)

	return conn, nil
}

func (p *corePoolImpl) Close() {

	p.Lock()
	defer p.Unlock()

	if p.isClosed {
		return
	}

	// Close all connections
	for _, closer := range p.closers {
		closer()
	}

	// Reset the pool state
	p.closers = nil
	p.connections = make(map[string]Connection)
	p.isClosed = true
}
