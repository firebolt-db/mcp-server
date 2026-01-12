package tool_search

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type SearchInput struct {
	Query string `json:"query"`
}

type SearchOutput struct {
	Result []mcp.Content `json:"result"`
}

type Search struct {
	topK                         int
	minScore                     float64
	includeChunks, includeScores bool

	baseURL string

	cl *http.Client
}

func (t *Search) Tool() *mcp.Tool {
	return &mcp.Tool{
		Name:  "firebolt_search",
		Title: "Firebolt Search",
		Description: "Returns search results from Firebolt knowledge base using RAG. " +
			"Use this tool to find answers to your questions about how to use Firebolt. ",
	}
}

func (t *Search) Register(s *mcp.Server) {
	{
		mcp.AddTool(s, t.Tool(), t.Handler())
	}
}

func (t *Search) Handler() mcp.ToolHandlerFor[SearchInput, *SearchOutput] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input SearchInput) (*mcp.CallToolResult, *SearchOutput, error) {

		return nil, nil, nil
	}
}

func NewSearch(ctx context.Context, cfg Config) (*Search, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config for search tool: %w", err)
	}

	// Initialize OAuth2 client credentials config
	authConfig := &clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     cfg.TokenURL,
		AuthStyle:    oauth2.AuthStyleInParams,
		EndpointParams: url.Values{
			"audience": []string{"https://api.firebolt.io"},
		},
	}

	return &Search{
		topK:          cfg.TopK,
		minScore:      cfg.MinScore,
		includeChunks: cfg.IncludeChunks,
		includeScores: cfg.IncludeScores,
		baseURL:       cfg.BaseURL,

		cl: authConfig.Client(ctx),
	}, nil
}
