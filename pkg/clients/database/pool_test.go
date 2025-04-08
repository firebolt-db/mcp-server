package database_test

import (
	"log/slog"
	"sync/atomic"
	"testing"

	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/database/databasemock"
)

func TestNewPool(t *testing.T) {

	t.Run("create new pool", func(t *testing.T) {
		pool, closer := database.NewPool(slogt.New(t), "", "")
		require.NotNil(t, pool, "Pool should not be nil")
		require.NotNil(t, closer, "Closer function should not be nil")
		closer()
	})

	t.Run("create pool with custom connection factory", func(t *testing.T) {

		factoryCalled := false
		customFactory := func(logger *slog.Logger, params database.DSNProvider) (database.ConnectionCloser, func(), error) {
			factoryCalled = true
			conn := databasemock.NewConnectionMock()
			return conn, conn.Close, nil
		}

		pool, closer := database.NewPoolWithConnectionFactory(slogt.New(t), "", "", customFactory)
		require.NotNil(t, pool, "Pool should not be nil")
		require.NotNil(t, closer, "Closer function should not be nil")

		// Get a connection to trigger the factory
		conn, err := pool.GetConnection(database.PoolParams{
			AccountName: "test-account",
		})
		require.NoError(t, err, "GetConnection should not return an error")
		require.NotNil(t, conn, "Connection should not be nil")
		assert.True(t, factoryCalled, "Custom connection factory should be called")

		closer()
	})
}

func TestPoolGetConnection(t *testing.T) {

	t.Run("get connection", func(t *testing.T) {

		factoryCallCount := 0
		customFactory := func(logger *slog.Logger, params database.DSNProvider) (database.ConnectionCloser, func(), error) {
			factoryCallCount++
			conn := databasemock.NewConnectionMock()
			return conn, conn.Close, nil
		}

		pool, closer := database.NewPoolWithConnectionFactory(slogt.New(t), "", "", customFactory)
		t.Cleanup(closer)

		params := database.PoolParams{
			AccountName: "test-account",
		}

		// First request should create a new connection
		conn1, err := pool.GetConnection(params)
		require.NoError(t, err, "GetConnection should not return an error")
		require.NotNil(t, conn1, "Connection should not be nil")
		assert.Equal(t, 1, factoryCallCount, "Connection factory should be called once")

		// Second request with same params should reuse the connection
		conn2, err := pool.GetConnection(params)
		require.NoError(t, err, "GetConnection should not return an error")
		require.NotNil(t, conn2, "Connection should not be nil")
		assert.Equal(t, 1, factoryCallCount, "Connection factory should still be called only once")

		// Connection objects should be the same instance
		assert.Same(t, conn1, conn2, "Both connections should be the same instance")
	})

	t.Run("get different connections", func(t *testing.T) {

		factoryCallCount := 0
		customFactory := func(logger *slog.Logger, params database.DSNProvider) (database.ConnectionCloser, func(), error) {
			factoryCallCount++
			conn := databasemock.NewConnectionMock()
			return conn, conn.Close, nil
		}

		pool, poolCloser := database.NewPoolWithConnectionFactory(slogt.New(t), "", "", customFactory)
		t.Cleanup(poolCloser)

		// Create connections with different parameters
		params1 := database.PoolParams{
			AccountName: "account1",
		}

		params2 := database.PoolParams{
			AccountName: "account2",
		}

		conn1, err := pool.GetConnection(params1)
		require.NoError(t, err, "GetConnection with params1 should not return an error")

		conn2, err := pool.GetConnection(params2)
		require.NoError(t, err, "GetConnection with params2 should not return an error")

		assert.NotSame(t, conn1, conn2, "Connections with different params should be different instances")
		assert.Equal(t, 2, factoryCallCount, "Connection factory should create 2 connections")
	})
}

func TestPoolClose(t *testing.T) {

	t.Run("close all connections", func(t *testing.T) {

		closeCount := 0
		customFactory := func(logger *slog.Logger, params database.DSNProvider) (database.ConnectionCloser, func(), error) {
			conn := databasemock.NewConnectionMock().WithCloseFunc(func() {
				closeCount++
			})
			return conn, conn.Close, nil
		}

		pool, poolCloser := database.NewPoolWithConnectionFactory(slogt.New(t), "", "", customFactory)

		// Create several connections
		params1 := database.PoolParams{AccountName: "account1"}
		params2 := database.PoolParams{AccountName: "account2"}

		_, err := pool.GetConnection(params1)
		require.NoError(t, err)

		_, err = pool.GetConnection(params2)
		require.NoError(t, err)

		// Close the pool - should close all connections
		poolCloser()
		assert.Equal(t, 2, closeCount, "All connections should be closed")

		// Verify pool is closed by trying to get a connection again
		conn3, err := pool.GetConnection(params1)
		require.Error(t, err, "GetConnection should return error after pool is closed")
		assert.Nil(t, conn3, "Connection should be nil after pool is closed")
		assert.ErrorIs(t, err, database.ErrPoolClosed, "Error should be ErrPoolClosed")
	})

	t.Run("multiple close calls", func(t *testing.T) {

		closeCount := 0
		customFactory := func(logger *slog.Logger, params database.DSNProvider) (database.ConnectionCloser, func(), error) {
			conn := databasemock.NewConnectionMock().WithCloseFunc(func() {
				closeCount++
			})
			return conn, conn.Close, nil
		}

		pool, poolCloser := database.NewPoolWithConnectionFactory(slogt.New(t), "", "", customFactory)

		// Create a connection
		params := database.PoolParams{AccountName: "account"}
		_, err := pool.GetConnection(params)
		require.NoError(t, err)

		// Close the pool multiple times - should be safe
		poolCloser()
		assert.Equal(t, 1, closeCount, "Connection should be closed once")

		poolCloser() // Second close call
		assert.Equal(t, 1, closeCount, "Connections should not be closed again")

		// Getting a new connection after closing should fail
		conn, err := pool.GetConnection(params)
		require.Error(t, err, "GetConnection should return error after pool is closed")
		assert.Nil(t, conn, "Connection should be nil after pool is closed")
		assert.ErrorIs(t, err, database.ErrPoolClosed, "Error should be ErrPoolClosed")
	})
}

func TestPoolConcurrency(t *testing.T) {

	t.Run("concurrent connection requests", func(t *testing.T) {

		var factoryCallCount atomic.Int32

		customFactory := func(logger *slog.Logger, params database.DSNProvider) (database.ConnectionCloser, func(), error) {
			factoryCallCount.Add(1)

			conn := databasemock.NewConnectionMock()
			return conn, conn.Close, nil
		}

		pool, closer := database.NewPoolWithConnectionFactory(slogt.New(t), "", "", customFactory)
		t.Cleanup(closer)

		params := database.PoolParams{
			AccountName: "test-account",
		}

		// Run concurrent goroutines all requesting the same connection parameters
		const numGoroutines = 100
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				conn, err := pool.GetConnection(params)
				assert.NoError(t, err)
				assert.NotNil(t, conn)
				done <- true
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			<-done
		}

		// Should only create one connection despite concurrent requests
		assert.Equal(t, int32(1), factoryCallCount.Load(),
			"Connection factory should be called exactly once despite concurrent requests")
	})
}
