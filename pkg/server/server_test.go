package server_test

import (
	"context"
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/server"
)

// TestNewServer checks that NewServer registers tools, prompts, and resource templates correctly.
func TestNewServer(t *testing.T) {
	// Create test logger
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	// Setup mocks
	mockTool1 := new(mockTool)
	mockTool1.On("Tool").Return(mcp.Tool{
		Name:        "testTool1",
		Description: "Test tool 1",
	})

	mockPrompt1 := new(mockPrompt)
	mockPrompt1.On("Prompt").Return(mcp.Prompt{
		Name:        "testPrompt1",
		Description: "Test prompt 1",
	})

	mockResource1 := new(mockResourceTemplate)
	mockResource1.On("ResourceTemplate").Return(
		mcp.NewResourceTemplate(
			"test",
			"testResource1",
			mcp.WithTemplateDescription("Test resource 1"),
		),
	)

	// Create server
	srv := server.NewServer(
		logger,
		"0.0.0",
		"stdio",
		"localhost:8080",
		[]server.Tool{mockTool1},
		[]server.Prompt{mockPrompt1},
		[]server.ResourceTemplate{mockResource1},
	)

	// Assertions
	require.NotNil(t, srv, "Server should not be nil")

	// Verify that mocks were called as expected
	mockTool1.AssertCalled(t, "Tool")
	mockPrompt1.AssertCalled(t, "Prompt")
	mockResource1.AssertCalled(t, "ResourceTemplate")
}

// TestServeUnsupportedTransport checks that Serve returns an error for unsupported transports.
func TestServeUnsupportedTransport(t *testing.T) {

	// Create server with unsupported transport
	srv := server.NewServer(slogt.New(t), "0.0.0", "test", "", nil, nil, nil)

	// Call Serve and check error
	err := srv.Serve(t.Context())
	assert.Error(t, err, "Serve should return an error for unsupported transport")
	assert.True(t, strings.Contains(err.Error(), "unsupported transport type"),
		"Error message should mention unsupported transport type")
}

// mockTool is a mock implementation of the Tool interface for testing.
type mockTool struct {
	mock.Mock
}

func (m *mockTool) Tool() mcp.Tool {
	args := m.Called()
	return args.Get(0).(mcp.Tool)
}

func (m *mockTool) Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mcp.CallToolResult), args.Error(1)
}

// mockPrompt is a mock implementation of the Prompt interface for testing.
type mockPrompt struct {
	mock.Mock
}

func (m *mockPrompt) Prompt() mcp.Prompt {
	args := m.Called()
	return args.Get(0).(mcp.Prompt)
}

func (m *mockPrompt) Handler(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mcp.GetPromptResult), args.Error(1)
}

// mockResourceTemplate is a mock implementation of the ResourceTemplate interface for testing.
type mockResourceTemplate struct {
	mock.Mock
}

func (m *mockResourceTemplate) ResourceTemplate() mcp.ResourceTemplate {
	args := m.Called()
	return args.Get(0).(mcp.ResourceTemplate)
}

func (m *mockResourceTemplate) Handler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]mcp.ResourceContents), args.Error(1)
}
