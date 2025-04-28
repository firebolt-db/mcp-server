package database

import (
	"fmt"
	"log/slog"
	"sync"
)

// ErrPoolClosed is returned when an operation is attempted on a closed pool.
var ErrPoolClosed = fmt.Errorf("pool is closed")

// Pool implements a thread-safe connection pool for Firebolt database connections.
//
// This MCP server is designed to provide access to multiple Firebolt accounts
// and databases simultaneously, which the current identity has access to.
// Since the Firebolt Go SDK creates connections to a single account only,
// we need to maintain a pool of connections to different accounts.
//
// The implementation creates a pool of connections for different combinations of:
// - Accounts
// - Databases
// - Engines
//
// This approach helps LLMs switch between different databases and engines
// with fewer mistakes, trading some connection efficiency for improved reliability.
type Pool interface {

	// GetConnection returns a connection for the specified parameters.
	// It first checks if a matching connection exists in the pool.
	// If not found, it creates a new connection, adds it to the pool, and returns it.
	GetConnection(params PoolParams) (Connection, error)

	// Close closes all connections in the pool and releases associated resources.
	Close()
}

// NewPool creates a new connection pool with the provided logger.
// It returns the pool implementation and a function to close all connections.
func NewPool(logger *slog.Logger, clientID, clientSecret string) (Pool, func()) {
	return NewPoolWithConnectionFactory(logger, clientID, clientSecret, NewConnection)
}

// NewConnectionFunc defines a function type for creating new connections.
type NewConnectionFunc func(*slog.Logger, DSNProvider) (ConnectionCloser, func(), error)

// NewPoolWithConnectionFactory creates a new connection pool with a custom connection factory.
// It allows you to provide a custom function to create connections which can be useful for testing.
func NewPoolWithConnectionFactory(
	logger *slog.Logger,
	clientID, clientSecret string,
	newConnectionFunc NewConnectionFunc,
) (Pool, func()) {

	pool := &poolImpl{
		logger:            logger,
		connections:       make(map[string]Connection),
		newConnectionFunc: newConnectionFunc,
		clientID:          clientID,
		clientSecret:      clientSecret,
	}

	return pool, pool.Close
}

type poolImpl struct {
	sync.Mutex
	isClosed          bool
	logger            *slog.Logger
	closers           []func()
	connections       map[string]Connection
	newConnectionFunc NewConnectionFunc
	clientID          string
	clientSecret      string
}

func (p *poolImpl) GetConnection(params PoolParams) (Connection, error) {

	p.Lock()
	defer p.Unlock()

	connectionParams := ConnectionParams{
		ClientID:     p.clientID,
		ClientSecret: p.clientSecret,
		AccountName:  params.AccountName,
		DatabaseName: params.DatabaseName,
		EngineName:   params.EngineName,
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

func (p *poolImpl) Close() {

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
