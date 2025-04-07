package discovery

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	authTokenEndpoint  = "/oauth/token"
	myAccountsEndpoint = "/myAccounts"
)

// Client represents a Firebolt API client that can be used to discover a list of accounts accessible by the identity
// on behalf of which the MCP server runs.
// The standard Firebolt Go SDK requires an account to be specified in the connection string.
// This client allows the MCP to work with multiple accounts within an organization and switch between them as needed.
type Client interface {

	// ListAccounts retrieves a list of accounts that the current identity has access to.
	// It returns a slice of Account or an error if the operation fails.
	ListAccounts(ctx context.Context) ([]Account, error)
}

// NewClient creates a new instance of the Firebolt API client.
// It initializes the client with the provided OAuth2 credentials and base URLs for authentication and API access.
// Returns an instance of Client or an error if the initialization fails.
func NewClient(
	ctx context.Context,
	logger *slog.Logger,
	clientID, clientSecret,
	authBaseURL, apiBaseURL string,
) (Client, error) {

	// Prepare config for OAuth2 client credentials authentication
	tokenURL, err := url.JoinPath(authBaseURL, authTokenEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not create token URL: %w", err)
	}
	authConfig := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		AuthStyle:    oauth2.AuthStyleInParams,
		EndpointParams: url.Values{
			"audience": []string{"https://api.firebolt.io"},
		},
	}

	// Return a new client implementation with the configured HTTP client
	return &clientImpl{
		apiBaseURL: apiBaseURL,
		logger:     logger,
		httpClient: authConfig.Client(ctx),
	}, nil
}

type clientImpl struct {
	apiBaseURL string
	logger     *slog.Logger
	httpClient *http.Client
}

func (c *clientImpl) ListAccounts(ctx context.Context) ([]Account, error) {

	// Construct the URL for the myAccounts endpoint
	myAccountsURL, err := url.JoinPath(c.apiBaseURL, myAccountsEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not create list accounts URL: %w", err)
	}

	// Create a new HTTP request with the constructed URL
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, myAccountsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Perform the HTTP request
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer func() {

		// Ensure the response body is fully read and closed
		if _, err := io.Copy(io.Discard, response.Body); err != nil {
			c.logger.Warn("failed to discard HTTP response body", "error", err)
		}
		if err := response.Body.Close(); err != nil {
			c.logger.Warn("failed to close HTTP response body", "error", err)
		}
	}()

	// Check if the response status code is OK
	if response.StatusCode != http.StatusOK {
		var responseText string
		responseData, err := io.ReadAll(response.Body)
		if err == nil {
			responseText = string(responseData)
		}
		return nil, fmt.Errorf(
			"unexpected response status code returned from Firebolt API: code - %d, responseText - %s",
			response.StatusCode, responseText,
		)
	}

	// Decode the response body JSON into a myAccountsResponse struct
	var result myAccountsResponse
	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body JSON: %w", err)
	}

	// Map the account names from the response
	return result.Accounts, nil
}

// Account represents a Firebolt account with its name and region.
type Account struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

type myAccountsResponse struct {
	Accounts []Account `json:"accounts"`
}
