package discovery_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"

	"github.com/firebolt-db/mcp-server/pkg/clients/discovery"
)

const (
	authTokenEndpoint  = "/oauth/token"
	myAccountsEndpoint = "/myAccounts"
)

type myAccountsResponse struct {
	Accounts []discovery.Account `json:"accounts"`
}

func TestListAccounts(t *testing.T) {

	t.Run("successful response", func(t *testing.T) {

		// Mock server
		mockServer := newMockServer(t, nil, nil)
		t.Cleanup(mockServer.Close)

		// Create client
		client, err := discovery.NewClient(
			context.WithValue(t.Context(), oauth2.HTTPClient, mockServer.Client()),
			slogt.New(t),
			"test-client-id",
			"test-client-secret",
			mockServer.URL,
			mockServer.URL,
		)
		require.NoError(t, err)

		// Call ListAccounts
		accounts, err := client.ListAccounts(t.Context())
		require.NoError(t, err)

		// Verify result
		expectedAccounts := []discovery.Account{
			{Name: "Account 1", Region: "us-east-1"},
			{Name: "Account 2", Region: "eu-west-1"},
		}
		assert.Equal(t, expectedAccounts, accounts)
	})

	t.Run("error response", func(t *testing.T) {

		// Mock server with error response for myAccounts endpoint
		errorHandler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(`{"error": "server error"}`))
			require.NoError(t, err)
		}
		mockServer := newMockServer(t, nil, errorHandler)
		t.Cleanup(mockServer.Close)

		// Create client
		client, err := discovery.NewClient(
			context.WithValue(t.Context(), oauth2.HTTPClient, mockServer.Client()),
			slogt.New(t),
			"test-client-id",
			"test-client-secret",
			mockServer.URL,
			mockServer.URL,
		)
		require.NoError(t, err)

		// Call ListAccounts - should fail
		accounts, err := client.ListAccounts(t.Context())
		assert.Error(t, err)
		assert.Nil(t, accounts)
		assert.Contains(t, err.Error(), "unexpected response status code")
	})

	t.Run("authentication error", func(t *testing.T) {

		// Mock server with authentication error
		authErrorHandler := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(`{"error": "invalid_client", "error_description": "Client authentication failed"}`))
			require.NoError(t, err)
		}
		mockServer := newMockServer(t, authErrorHandler, nil)
		t.Cleanup(mockServer.Close)

		// Create client
		client, err := discovery.NewClient(
			context.WithValue(t.Context(), oauth2.HTTPClient, mockServer.Client()),
			slogt.New(t),
			"invalid-client-id",
			"invalid-client-secret",
			mockServer.URL,
			mockServer.URL,
		)
		require.NoError(t, err)

		// Call ListAccounts - should fail due to authentication error
		accounts, err := client.ListAccounts(t.Context())
		assert.Error(t, err)
		assert.Nil(t, accounts)
		// oauth2 could have various error messages, so we don't check the exact error content
	})

	t.Run("invalid JSON response", func(t *testing.T) {

		// Mock server with invalid JSON response
		invalidJSONHandler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"accounts": [{"name": "Account 1"`)) // Incomplete JSON
			require.NoError(t, err)
		}
		mockServer := newMockServer(t, nil, invalidJSONHandler)
		t.Cleanup(mockServer.Close)

		// Create client
		client, err := discovery.NewClient(
			context.WithValue(t.Context(), oauth2.HTTPClient, mockServer.Client()),
			slogt.New(t),
			"test-client-id",
			"test-client-secret",
			mockServer.URL,
			mockServer.URL,
		)
		require.NoError(t, err)

		// Call ListAccounts - should fail due to invalid JSON
		accounts, err := client.ListAccounts(t.Context())
		assert.Error(t, err)
		assert.Nil(t, accounts)
		assert.Contains(t, err.Error(), "failed to decode response body JSON")
	})

	t.Run("context cancellation", func(t *testing.T) {

		// Mock server with delay
		slowHandler := func(w http.ResponseWriter, r *http.Request) {
			// This handler deliberately does nothing and doesn't respond,
			// which will cause the request to wait until context cancellation
			select {
			case <-r.Context().Done():
				// Context was cancelled, we can return now
				return
			}
		}
		mockServer := newMockServer(t, nil, slowHandler)
		t.Cleanup(mockServer.Close)

		// Create client
		client, err := discovery.NewClient(
			context.WithValue(t.Context(), oauth2.HTTPClient, mockServer.Client()),
			slogt.New(t),
			"test-client-id",
			"test-client-secret",
			mockServer.URL,
			mockServer.URL,
		)
		require.NoError(t, err)

		// Create a context that is already cancelled
		ctx, cancel := context.WithCancel(t.Context())
		cancel() // Cancel immediately

		// Call ListAccounts with cancelled context
		accounts, err := client.ListAccounts(ctx)
		assert.Error(t, err)
		assert.Nil(t, accounts)
		assert.Contains(t, err.Error(), "context canceled")
	})

	t.Run("server not available", func(t *testing.T) {

		// Create a server and immediately close it to ensure the URL is unreachable
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		mockServer.Close() // Close immediately to make server unavailable

		// Create client with URL that points to the closed server
		client, err := discovery.NewClient(
			t.Context(),
			slogt.New(t),
			"test-client-id",
			"test-client-secret",
			mockServer.URL,
			mockServer.URL,
		)
		require.NoError(t, err)

		// Call ListAccounts - should fail because server is not available
		accounts, err := client.ListAccounts(t.Context())
		assert.Error(t, err)
		assert.Nil(t, accounts)
		// Error message may vary based on system/network, so we don't check exact text
	})
}

func TestNewClient(t *testing.T) {

	t.Run("valid configuration", func(t *testing.T) {
		client, err := discovery.NewClient(
			t.Context(),
			slogt.New(t),
			"client-id",
			"client-secret",
			"https://auth.example.com",
			"https://api.example.com",
		)
		require.NoError(t, err)
		assert.NotNil(t, client)
	})

	t.Run("invalid auth URL", func(t *testing.T) {
		client, err := discovery.NewClient(
			t.Context(),
			slogt.New(t),
			"client-id",
			"client-secret",
			"://invalid-url", // Invalid URL
			"https://api.example.com",
		)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "could not create token URL")
	})
}

func newMockServer(
	t *testing.T,
	customAuthHandler http.HandlerFunc,
	customAccountsHandler http.HandlerFunc,
) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// /auth/token endpoint
		if r.Method == http.MethodPost && r.URL.Path == authTokenEndpoint {

			if customAuthHandler != nil {
				customAuthHandler(w, r)
				return
			}

			// Return a successful response
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(`{"access_token":"test-token","token_type":"bearer","expires_in":3600}`))
			require.NoError(t, err)
			return
		}

		// /myAccounts endpoint
		if r.Method == http.MethodGet && r.URL.Path == myAccountsEndpoint {

			if customAccountsHandler != nil {
				customAccountsHandler(w, r)
				return
			}

			// Return a successful response
			response := myAccountsResponse{
				Accounts: []discovery.Account{
					{Name: "Account 1", Region: "us-east-1"},
					{Name: "Account 2", Region: "eu-west-1"},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			require.NoError(t, json.NewEncoder(w).Encode(response))
			return
		}

		http.NotFound(w, r)
	}))
}
