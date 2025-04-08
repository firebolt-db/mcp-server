// Package server provides the core MCP server implementation for the Firebolt MCP.
// It defines interfaces for tools, prompts, and resource templates that can be registered
// with the server, as well as methods for serving the MCP over different transports.
package server

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"

	"github.com/firebolt-db/mcp-server/pkg/version"
)

// Server represents the interface for the MCP server functionality.
// It defines the method to start the server and handle MCP requests.
type Server interface {
	// Serve starts the server with the provided context and handles MCP requests
	// until the context is canceled or an error occurs.
	Serve(ctx context.Context) error
}

// Tool represents a callable tool that can be registered with the MCP server.
// It provides methods to define the tool metadata and handle tool calls.
type Tool interface {
	// Tool returns the MCP tool definition with its name, description, parameters, etc.
	Tool() mcp.Tool
	// Handler processes tool call requests and returns results.
	Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Prompt represents a prompt that can be registered with the MCP server.
// It provides methods to define the prompt metadata and handle prompt requests.
type Prompt interface {
	// Prompt returns the MCP prompt definition.
	Prompt() mcp.Prompt
	// Handler processes prompt requests and returns results.
	Handler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error)
}

// ResourceTemplate represents a resource template that can be registered with the MCP server.
// It provides methods to define the resource template metadata and handle resource requests.
type ResourceTemplate interface {
	// ResourceTemplate returns the MCP resource template definition.
	ResourceTemplate() mcp.ResourceTemplate
	// Handler processes resource read requests and returns resource contents.
	Handler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error)
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
	transport string,
	transportSSEAddress string,
	tools []Tool,
	prompts []Prompt,
	resourceTemplates []ResourceTemplate,
) Server {

	// Configure logging hooks to track tool calls and errors
	hooks := &mcpserver.Hooks{}
	hooks.AddBeforeCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest) {
		logger.DebugContext(
			ctx,
			"received tool call request",
			slog.Any("id", id),
			slog.String("tool", message.Params.Name),
			slog.Any("arguments", message.Params.Arguments),
		)
	})
	hooks.AddAfterCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		logger.InfoContext(
			ctx,
			"tool call finished",
			slog.Any("id", id),
			slog.String("tool", message.Params.Name),
			slog.Any("any", result.Result),
		)
	})
	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		logger.ErrorContext(
			ctx,
			"error occurred",
			slog.Any("id", id),
			slog.String("method", string(method)),
			slog.Any("message", message),
			slog.String("error", err.Error()),
		)
	})

	// Initialize the MCP server with Firebolt-specific configuration
	mcpSrv := mcpserver.NewMCPServer(
		"Firebolt MCP Server",
		version.GetFullVersion(),
		mcpserver.WithInstructions(`
			This MCP makes you a Firebolt cloud data warehouse expert with access to specialized tools and resources.
			You can assist with SQL queries, data modeling, performance optimization, and analytics for Firebolt.
			
			Use available tools and resources to:
			- Access Firebolt documentation for reference
			- Execute SQL queries against Firebolt databases
		`),
		mcpserver.WithHooks(hooks),
		mcpserver.WithToolCapabilities(true),
		mcpserver.WithPromptCapabilities(false),
		mcpserver.WithResourceCapabilities(false, false),
	)

	// Register the tools, prompts, and resource templates with the server
	mcpSrv.AddTools(transform(tools, func(i Tool) mcpserver.ServerTool {
		return mcpserver.ServerTool{
			Tool:    i.Tool(),
			Handler: i.Handler,
		}
	})...)
	for _, prompt := range prompts {
		mcpSrv.AddPrompt(prompt.Prompt(), prompt.Handler)
	}
	for _, resourceTemplate := range resourceTemplates {
		mcpSrv.AddResourceTemplate(resourceTemplate.ResourceTemplate(), resourceTemplate.Handler)
	}

	// Initialize the server implementation with the configured MCP server
	s := &serverImpl{
		logger:              logger,
		inner:               mcpSrv,
		transport:           transport,
		transportSSEAddress: transportSSEAddress,
	}

	return s
}

// serverImpl is the implementation of the Server interface.
// It wraps an MCP server and provides methods to serve it over different transports.
type serverImpl struct {
	logger              *slog.Logger
	inner               *mcpserver.MCPServer
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

	srv := mcpserver.NewStdioServer(s.inner)
	srv.SetErrorLogger(log.New(os.Stderr, "", log.LstdFlags))

	return srv.Listen(ctx, os.Stdin, os.Stdout)
}

// serveSSE starts the server using the SSE transport.
// It listens for HTTP requests on the configured address and serves MCP requests over SSE.
// The server will shut down gracefully when the context is canceled.
func (s *serverImpl) serveSSE(ctx context.Context) error {
	s.logger.Info("Using sse transport", "listen_address", s.transportSSEAddress)

	srv := mcpserver.NewSSEServer(s.inner)

	// Setup graceful shutdown when context is canceled
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Warn("error shutting down server gracefully", "error", err)
		}
	}()

	return srv.Start(s.transportSSEAddress)
}

// transform is a utility function that applies a transformation function to each element
// of the input slice and returns a new slice containing the transformed elements.
// It's a generic function that works with any input and output types.
func transform[I, O any](inputs []I, transformer func(I) O) []O {
	outputs := make([]O, len(inputs))
	for i, input := range inputs {
		outputs[i] = transformer(input)
	}
	return outputs
}
