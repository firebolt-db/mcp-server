package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Server represents the interface for the MCP server functionality.
// It defines the method to start the server and handle MCP requests.
type Server interface {
	// Serve starts the server with the provided context and handles MCP requests
	// until the context is canceled or an error occurs.
	Serve(ctx context.Context) error
}

// Tool represents a callable tool that can be registered with the MCP server.
// It provides methods to define the tool metadata and register it in mcp server.
type Tool interface {
	// Tool returns the MCP tool definition.
	Tool() *mcp.Tool
	// Register adds the tool to the server using its specific generic types.
	Register(s *mcp.Server)
}

// Prompt represents a prompt that can be registered with the MCP server.
// It provides methods to define the prompt metadata and handle prompt requests.
type Prompt interface {
	// Prompt returns the MCP prompt definition.
	Prompt() *mcp.Prompt
	// Handler processes prompt requests and returns results.
	Handler(ctx context.Context, request *mcp.GetPromptRequest) (*mcp.GetPromptResult, error)
}

// ResourceTemplate represents a resource template that can be registered with the MCP server.
// It provides methods to define the resource template metadata and handle resource requests.
type ResourceTemplate interface {
	// ResourceTemplate returns the MCP resource template definition.
	ResourceTemplate() *mcp.ResourceTemplate
	// Handler processes resource read requests and returns resource contents.
	Handler(ctx context.Context, request *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error)
}

// NewServer creates a new MCP server with the provided configuration.
// It registers the provided tools, prompts, and resource templates with the server.
//
// Parameters:
//   - logger: The structured logger to use for logging.
//   - transport: The transport type to use ("stdio" or "sse").
//   - transportSSEAddress: The address to listen on when using SSE transport.
//   - tools: The list of tools to register with the server.
//   - prompts: The list of prompts to register with the server.
//   - resourceTemplates: The list of resource templates to register with the server.
//
// Returns a Server that can be used to start handling MCP requests.
func NewServer(
	logger *slog.Logger,
	version string,
	transport string,
	transportSSEAddress string,
	tools []Tool,
	prompts []Prompt,
	resourceTemplates []ResourceTemplate,
) Server {

	mcpSrv := mcp.NewServer(
		&mcp.Implementation{
			Name:       "Firebolt MCP Server",
			Title:      "Firebolt MCP Server",
			Version:    version,
			WebsiteURL: "https://firebolt.io",
		},
		&mcp.ServerOptions{
			Instructions: `This MCP makes you a Firebolt cloud data warehouse expert with access to specialized tools and resources.
			You can assist with SQL queries, data modeling, performance optimization, and analytics for Firebolt.
			
			Use available tools and resources to:
			- Access Firebolt documentation for reference
			- Execute SQL queries against Firebolt databases`,
		},
	)

	// register tool call logging middleware
	mcpSrv.AddReceivingMiddleware(logging(logger))

	for _, tool := range tools {
		tool.Register(mcpSrv)
	}

	for _, prompt := range prompts {
		mcpSrv.AddPrompt(prompt.Prompt(), prompt.Handler)
	}

	for _, resourceTemplate := range resourceTemplates {
		mcpSrv.AddResourceTemplate(resourceTemplate.ResourceTemplate(), resourceTemplate.Handler)
	}

	return &serverImpl{
		logger:              logger,
		inner:               mcpSrv,
		transport:           transport,
		transportSSEAddress: transportSSEAddress,
	}
}

// serverImpl is the implementation of the Server interface.
// It wraps an MCP server and provides methods to serve it over different transports.
type serverImpl struct {
	logger              *slog.Logger
	inner               *mcp.Server
	transport           string
	transportSSEAddress string
}

// Serve starts the server with the provided context and handles MCP requests
// until the context is canceled or an error occurs.
// It selects the appropriate transport (stdio or SSE) based on the server configuration.
func (s *serverImpl) Serve(ctx context.Context) error {
	switch s.transport {
	case "stdio":
		return s.serveStdio(ctx)
	case "sse":
		return s.serveSSE(ctx)
	default:
		return errors.New("unsupported transport type: " + s.transport)
	}
}

// serveStdio starts the server using the stdio transport.
// It reads MCP requests from stdin and writes responses to stdout.
func (s *serverImpl) serveStdio(ctx context.Context) error {
	s.logger.Info("Using stdio transport")

	return s.inner.Run(ctx, &mcp.StdioTransport{})
}

func (s *serverImpl) serveSSE(ctx context.Context) error {
	s.logger.Info("Using sse transport", "listen_address", s.transportSSEAddress)

	handler := mcp.NewSSEHandler(func(request *http.Request) *mcp.Server {
		return s.inner
	}, nil)

	// Configure the HTTP server
	srv := &http.Server{
		Addr:    s.transportSSEAddress,
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Warn("error shutting down server gracefully", "error", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
