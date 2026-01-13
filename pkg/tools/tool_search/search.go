package tool_search

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// defaultTopK is the default number of documents to retrieve
	defaultTopK = 10

	// defaultScoreThreshold is the default minimum similarity score (-1.0 to 1.0) required for a document to be included in results
	defaultScoreThreshold = 0.4
)

type Input struct {
	Query string `json:"query"`
}

type Output struct {
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

func (t *Search) Handler() mcp.ToolHandlerFor[Input, *Output] {
	return func(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, *Output, error) {
		if input.Query == "" {
			return nil, nil, fmt.Errorf("bad request: query parameter is required")
		}

		params := url.Values{}
		params.Add("topK", fmt.Sprintf("%d", t.topK))
		params.Add("minScore", fmt.Sprintf("%f", t.minScore))
		params.Add("includeChunks", fmt.Sprintf("%t", t.includeChunks))
		params.Add("includeScores", fmt.Sprintf("%t", t.includeScores))

		// run the search query
		r, err := http.NewRequestWithContext(ctx,
			http.MethodGet,
			t.baseURL+"/docs/v1/search/"+url.QueryEscape(input.Query),
			strings.NewReader(params.Encode()),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create search request: %w", err)
		}

		resp, err := t.cl.Do(r)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to make search request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, nil, fmt.Errorf("failed to make search request, status: %d", resp.StatusCode)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read response body: %w", err)
		}

		out := &Output{
			Result: []mcp.Content{
				&mcp.TextContent{
					Text: string(data),
				},
			},
		}

		return &mcp.CallToolResult{}, out, nil
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
		topK:          defaultTopK,
		minScore:      defaultScoreThreshold,
		includeChunks: true, // include chunks by default
		includeScores: true, // include scores by default
		baseURL:       cfg.BaseURL,

		cl: authConfig.Client(ctx),
	}, nil
}
