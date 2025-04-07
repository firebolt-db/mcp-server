package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"sync/atomic"

	_ "github.com/firebolt-db/firebolt-go-sdk" // Register the Firebolt driver
)

// ErrConnectionClosed is returned when an operation is attempted on a closed connection.
var ErrConnectionClosed = fmt.Errorf("connection is closed")

// Connection defines the methods for interacting with Firebolt via SQL.
//
// It provides a standardized interface to execute queries against a Firebolt database instance
// while abstracting away the underlying implementation details of the connection.
type Connection interface {

	// Query executes a SQL query against Firebolt and returns the result.
	// The sql parameter can be any valid Firebolt SQL statement, including
	// SELECT, INSERT, UPDATE, DELETE, etc.
	Query(ctx context.Context, sql string, args ...any) ([]map[string]any, error)
}

// ConnectionCloser extends the Connection interface to include a method for closing the connection.
type ConnectionCloser interface {
	Connection

	// Close terminates the underlying database connection and releases
	// all associated resources.
	Close()
}

// NewConnection creates a new Firebolt connection using the provided parameters.
// It returns a ConnectionCloser interface, a function to close the connection, and an error
// if any occurred during the connection setup.
func NewConnection(
	logger *slog.Logger,
	params DSNProvider,
) (ConnectionCloser, func(), error) {

	db, err := sql.Open(params.DriverName(), params.DSN())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open Firebolt connection: %w", err)
	}

	conn := &connectionImpl{
		db:       db,
		logger:   logger,
		isClosed: &atomic.Bool{},
	}

	return conn, conn.Close, nil
}

type connectionImpl struct {
	db       *sql.DB
	logger   *slog.Logger
	isClosed *atomic.Bool
}

func (c *connectionImpl) Query(ctx context.Context, sql string, args ...any) ([]map[string]any, error) {

	// Check if the connection is closed
	if c.isClosed.Load() {
		return nil, ErrConnectionClosed
	}

	// Query the database
	rows, err := c.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			c.logger.Warn(
				"failed to close database result rows",
				"error", err,
				"sql", sql,
			)
		}
	}()

	// Retrieve column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Iterate over the rows and scan the results into a slice of maps
	var results []map[string]any
	for rows.Next() {

		// Create containers for the row data
		values := make([]any, len(columns))
		valuePtrs := make([]any, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan the row into the value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Map column names to their values
		rowMap := make(map[string]any)
		for i, col := range columns {
			rowMap[col] = values[i]
		}

		results = append(results, rowMap)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return results, nil
}

func (c *connectionImpl) Close() {

	if c.isClosed.Load() {
		return
	}

	if err := c.db.Close(); err != nil {
		c.logger.Warn("failed to close database connection", "error", err)
	}

	c.isClosed.Store(true)
}

// DSNProvider defines the interface for providing a Data Source Name (DSN).
// The only DSNProvider implementation is ConnectionParams, but this interface
// allows to write mock implementations for testing purposes.
type DSNProvider interface {

	// DriverName returns the name of the database driver.
	DriverName() string

	// DSN returns the Data Source Name (DSN) for the connection.
	DSN() string
}
