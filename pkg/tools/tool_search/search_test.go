package tool_search_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/tools/tool_search"
)

const (
	clientID     = "client_id"
	clientSecret = "client_secret"
)

func TestNewSearch(t *testing.T) {
	t.Parallel()

	search, err := tool_search.NewSearch(t.Context(), newConfig())
	require.NoError(t, err)
	require.NotNil(t, search)
}

func TestNewSearch_InvalidConfig(t *testing.T) {
	t.Parallel()

	_, err := tool_search.NewSearch(t.Context(), tool_search.Config{})
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid config for search tool")
}

func TestNewSearch_Tool(t *testing.T) {
	t.Parallel()

	search, err := tool_search.NewSearch(t.Context(), newConfig())
	require.NoError(t, err)
	require.NotNil(t, search)

	tool := search.Tool()
	require.Equal(t, "firebolt_docs_search", tool.Name)
	assert.Contains(t, tool.Description, "Returns search results from Firebolt knowledge base using RAG")
}

func TestSearch_Success(t *testing.T) {
	t.Parallel()

	accessTkn := "access_token"
	in := tool_search.Input{
		Query: "Search query",
	}
	ragResp := `{"foo":"bar"}`

	cfg := newConfig()
	// Auth server will return valid token
	authSrv := fakeAuthServer(t, accessTokenResponse(t, accessTkn, 3600))

	// set up RAG server. Verify query and access token
	ragSrv := fakeServer(t, func(w http.ResponseWriter, r *http.Request) {
		at := r.Header.Get("Authorization")
		require.Equal(t, "Bearer "+accessTkn, at)

		require.Equal(t, r.URL.Path, "/docs/v1/search/"+url.QueryEscape(in.Query))
		w.Write([]byte(ragResp))
	})

	cfg.TokenURL = authSrv.URL
	cfg.BaseURL = ragSrv.URL

	search, err := tool_search.NewSearch(t.Context(), cfg)
	require.NoError(t, err)

	h := search.Handler()

	result, out, err := h(t.Context(), &mcp.CallToolRequest{}, in)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, out)
	require.Len(t, out.Result, 1)
	require.Equal(t, ragResp, out.Result[0].(*mcp.TextContent).Text)
}

func TestSearch_AuthenticationFailure(t *testing.T) {
	t.Parallel()

	cfg := newConfig()

	// Auth server will return 401 Unauthorized
	authSrv := fakeAuthServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})

	cfg.TokenURL = authSrv.URL

	search, err := tool_search.NewSearch(t.Context(), cfg)
	require.NoError(t, err)

	h := search.Handler()

	in := tool_search.Input{
		Query: "Search query",
	}

	result, out, err := h(t.Context(), &mcp.CallToolRequest{}, in)
	require.Error(t, err)
	require.ErrorContains(t, err, "oauth2: cannot fetch token: 401 Unauthorized")
	require.Nil(t, result)
	require.Nil(t, out)
}

func TestSearch_RAGRequestFailure(t *testing.T) {
	t.Parallel()

	accessTkn := "access_token"
	in := tool_search.Input{
		Query: "Search query",
	}

	cfg := newConfig()
	// Auth server will return valid token
	authSrv := fakeAuthServer(t, accessTokenResponse(t, accessTkn, 3600))

	// set up RAG server. Verify query and access token
	ragSrv := fakeServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	cfg.TokenURL = authSrv.URL
	cfg.BaseURL = ragSrv.URL

	search, err := tool_search.NewSearch(t.Context(), cfg)
	require.NoError(t, err)

	h := search.Handler()

	result, out, err := h(t.Context(), &mcp.CallToolRequest{}, in)
	require.Error(t, err)
	require.ErrorContains(t, err, fmt.Sprintf("failed to make search request, status: %d", http.StatusNotFound))
	require.Nil(t, result)
	require.Nil(t, out)
}

func newConfig() tool_search.Config {
	return tool_search.Config{
		BaseURL:      "https://api.local.firebolt.io",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://id.local.firebolt.io/oauth/token",
	}
}

func fakeServer(t *testing.T, fn http.HandlerFunc) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fn(w, r)
		},
	))
}

func fakeAuthServer(t *testing.T, fn http.HandlerFunc) *httptest.Server {
	t.Helper()

	return fakeServer(t, func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, clientID, r.FormValue("client_id"))
		require.Equal(t, clientSecret, r.FormValue("client_secret"))
		require.Equal(t, "client_credentials", r.FormValue("grant_type"))
		require.Equal(t, "https://api.firebolt.io", r.FormValue("audience"))

		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	})
}

func accessTokenResponse(t *testing.T, accessToken string, expiresIn int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"access_token": accessToken,
			"token_type":   "Bearer",
			"expires_in":   expiresIn,
		}
		bytes, err := json.Marshal(response)
		require.NoError(t, err)

		_, err = w.Write(bytes)
		require.NoError(t, err)
	}
}
