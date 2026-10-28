package server_test

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/neilotoole/slogt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/clients/database/databasemock"
	"github.com/firebolt-db/mcp-server/pkg/prompts"
	"github.com/firebolt-db/mcp-server/pkg/server"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_query"
)

// TestNewServer checks that NewServer registers tools, prompts, and resource templates correctly.
func TestNewServer(t *testing.T) {
	// Create test logger
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	// Setup mocks
	mockTool1 := new(mockTool)
	mockTool1.On("Tool").Return(&mcp.Tool{
		Name:        "testTool1",
		Description: "Test tool 1",
	})

	mockPrompt1 := new(mockPrompt)
	mockPrompt1.On("Prompt").Return(&mcp.Prompt{
		Name:        "testPrompt1",
		Description: "Test prompt 1",
	})

	mockResource1 := new(mockResourceTemplate)
	mockResource1.On("ResourceTemplate").Return(
		&mcp.ResourceTemplate{
			URITemplate: "test",
			Name:        "testResource1",
			Description: "Test resource 1",
		},
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

// TestServeStreamableHTTP checks that a Streamable HTTP client can initialize a session
// with a bare POST, without opening an SSE stream first, and discover tools and prompts.
func TestServeStreamableHTTP(t *testing.T) {
	addr := freeAddr(t)
	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", addr,
		[]server.Tool{&funcTool{name: "test_tool"}},
		[]server.Prompt{prompts.NewFireboltExpert()},
		nil,
	)
	cancel, serveErr := serve(t, srv)

	session := connect(t, addr, true)
	assert.Equal(t, "Firebolt MCP Server", session.InitializeResult().ServerInfo.Name)

	toolList, err := session.ListTools(t.Context(), nil)
	require.NoError(t, err)
	require.Len(t, toolList.Tools, 1)
	assert.Equal(t, "test_tool", toolList.Tools[0].Name)

	promptList, err := session.ListPrompts(t.Context(), nil)
	require.NoError(t, err)
	require.Len(t, promptList.Prompts, 1)
	assert.Equal(t, "firebolt_expert", promptList.Prompts[0].Name)

	require.NoError(t, session.Close())

	cancel()
	require.NoError(t, recv(t, serveErr, "Serve to return"), "Serve should shut down cleanly")
}

// TestStreamableHTTPRejectsBrowserAttacks checks that requests a malicious web page can make
// are rejected before they reach MCP handlers, while other clients are accepted.
func TestStreamableHTTPRejectsBrowserAttacks(t *testing.T) {
	var calls atomic.Int32
	addr := freeAddr(t)
	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", addr,
		[]server.Tool{&funcTool{name: "test_tool", fn: func(context.Context) { calls.Add(1) }}},
		nil,
		nil,
	)
	cancel, serveErr := serve(t, srv)
	waitListening(t, addr)
	_, port, err := net.SplitHostPort(addr)
	require.NoError(t, err)

	// Initializes a session and calls a tool in one request, so it needs no session ID.
	batch := `[` + initializeRequest + `,` +
		`{"jsonrpc":"2.0","method":"notifications/initialized"},` +
		`{"jsonrpc":"2.0","id":2,"method":"tools/call","params":{"name":"test_tool","arguments":{}}}]`

	tests := map[string]struct {
		host       string
		header     http.Header
		body       string
		wantStatus int
	}{
		"cross-site simple request": {
			// Browsers send text/plain and form POSTs to other origins without a CORS preflight.
			header: http.Header{
				"Origin":         {"https://evil.example"},
				"Sec-Fetch-Site": {"cross-site"},
				"Content-Type":   {"text/plain"},
			},
			body:       batch,
			wantStatus: http.StatusUnsupportedMediaType,
		},
		"cross-site form request": {
			header: http.Header{
				"Origin":         {"https://evil.example"},
				"Sec-Fetch-Site": {"cross-site"},
				"Content-Type":   {"application/x-www-form-urlencoded"},
			},
			body:       batch,
			wantStatus: http.StatusUnsupportedMediaType,
		},
		"DNS rebinding": {
			// After rebinding, the attacker's page is same-origin with the server.
			host: "rebind.evil.example:" + port,
			header: http.Header{
				"Origin":       {"http://rebind.evil.example:" + port},
				"Content-Type": {"application/json"},
			},
			body:       batch,
			wantStatus: http.StatusForbidden,
		},
		"client with Accept parameters": {
			header: http.Header{
				"Content-Type": {"application/json"},
				"Accept":       {"application/json;q=1, text/event-stream;q=1"},
			},
			body:       initializeRequest,
			wantStatus: http.StatusOK,
		},
		"client with Content-Type charset": {
			header:     http.Header{"Content-Type": {"application/json; charset=utf-8"}},
			body:       initializeRequest,
			wantStatus: http.StatusOK,
		},
		"gateway forwarding browser headers": {
			header: http.Header{
				"Origin":         {"https://chat.example"},
				"Sec-Fetch-Site": {"cross-site"},
				"Content-Type":   {"application/json"},
			},
			body:       initializeRequest,
			wantStatus: http.StatusOK,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, "http://"+addr+"/mcp", strings.NewReader(tt.body))
			require.NoError(t, err)
			req.Header = tt.header
			if req.Header.Get("Accept") == "" {
				req.Header.Set("Accept", "application/json, text/event-stream")
			}
			if tt.host != "" {
				req.Host = tt.host
			}

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}

	t.Run("no CORS preflight grant", func(t *testing.T) {
		// Without a granted preflight, browsers do not send JSON POSTs to other origins.
		req, err := http.NewRequestWithContext(t.Context(), http.MethodOptions, "http://"+addr+"/mcp", nil)
		require.NoError(t, err)
		req.Header.Set("Origin", "https://evil.example")
		req.Header.Set("Access-Control-Request-Method", http.MethodPost)
		req.Header.Set("Access-Control-Request-Headers", "content-type")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())
		assert.Empty(t, resp.Header.Get("Access-Control-Allow-Origin"))
	})

	cancel()
	require.NoError(t, recv(t, serveErr, "Serve to return"))
	assert.Zero(t, calls.Load(), "rejected requests should not run tools")
}

// TestSSERejectsBrowserAttacks checks that the SSE transport rejects cross-site simple requests
// and DNS rebinding.
func TestSSERejectsBrowserAttacks(t *testing.T) {
	addr := freeAddr(t)
	srv := server.NewServer(slogt.New(t), "0.0.0", "sse", addr, nil, nil, nil)
	cancel, serveErr := serve(t, srv)
	waitListening(t, addr)
	_, port, err := net.SplitHostPort(addr)
	require.NoError(t, err)

	t.Run("cross-site simple request", func(t *testing.T) {
		req, err := http.NewRequestWithContext(t.Context(), http.MethodPost, "http://"+addr+"/?sessionid=x", strings.NewReader(initializeRequest))
		require.NoError(t, err)
		req.Header.Set("Origin", "https://evil.example")
		req.Header.Set("Content-Type", "text/plain")

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())
		assert.Equal(t, http.StatusUnsupportedMediaType, resp.StatusCode)
	})

	t.Run("DNS rebinding", func(t *testing.T) {
		req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://"+addr+"/sse", nil)
		require.NoError(t, err)
		req.Host = "rebind.evil.example:" + port

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		require.NoError(t, resp.Body.Close())
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	})

	cancel()
	require.NoError(t, recv(t, serveErr, "Serve to return"))
}

// TestServeWaitsForInFlightRequests checks that shutdown lets a running tool call finish,
// and returns as soon as it does.
func TestServeWaitsForInFlightRequests(t *testing.T) {
	started := make(chan struct{})
	release := make(chan struct{})
	addr := freeAddr(t)
	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", addr,
		[]server.Tool{&funcTool{name: "slow_tool", fn: func(context.Context) {
			close(started)
			<-release
		}}},
		nil,
		nil,
	)
	server.SetShutdownTimeout(srv, 3*time.Second)
	cancel, serveErr := serve(t, srv)
	// Without the standalone stream, the tool call is the only open request.
	session := connect(t, addr, false)

	callErr := make(chan error, 1)
	go func() {
		_, err := session.CallTool(t.Context(), &mcp.CallToolParams{Name: "slow_tool"})
		callErr <- err
	}()
	recv(t, started, "tool to start")

	cancel()
	select {
	case err := <-serveErr:
		t.Fatalf("Serve returned while a tool call was running: %v", err)
	case <-time.After(200 * time.Millisecond):
	}

	close(release)
	require.NoError(t, recv(t, callErr, "tool call to finish"), "the running tool call should complete")
	start := time.Now()
	require.NoError(t, recv(t, serveErr, "Serve to return"))
	assert.Less(t, time.Since(start), time.Second, "Serve should not wait for the shutdown timeout")
}

// TestServeClosesConnectionsAfterShutdownTimeout checks that a request still running when the
// shutdown timeout expires is cut off instead of keeping the server alive.
func TestServeClosesConnectionsAfterShutdownTimeout(t *testing.T) {
	started := make(chan struct{})
	release := make(chan struct{})
	t.Cleanup(func() { close(release) })
	addr := freeAddr(t)
	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", addr,
		[]server.Tool{&funcTool{name: "stuck_tool", fn: func(context.Context) {
			close(started)
			<-release
		}}},
		nil,
		nil,
	)
	server.SetShutdownTimeout(srv, 100*time.Millisecond)
	cancel, serveErr := serve(t, srv)
	session := connect(t, addr, true)

	callErr := make(chan error, 1)
	go func() {
		_, err := session.CallTool(t.Context(), &mcp.CallToolParams{Name: "stuck_tool"})
		callErr <- err
	}()
	recv(t, started, "tool to start")

	cancel()
	start := time.Now()
	require.NoError(t, recv(t, serveErr, "Serve to return"))
	assert.Less(t, time.Since(start), time.Second, "Serve should return soon after the shutdown timeout")
	require.Error(t, recv(t, callErr, "tool call to fail"), "the cut-off tool call should fail")
}

// TestServeCanceledContext checks that Serve returns cleanly when its context is already canceled.
func TestServeCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", freeAddr(t), nil, nil, nil)
	assert.NoError(t, srv.Serve(ctx))
}

// TestServeReturnsListenError checks that Serve reports a listen address that is already in use.
func TestServeReturnsListenError(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = l.Close() })

	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", l.Addr().String(), nil, nil, nil)
	assert.Error(t, srv.Serve(t.Context()))
}

const initializeRequest = `{"jsonrpc":"2.0","id":1,"method":"initialize",` +
	`"params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"test","version":"0.0.0"}}}`

// freeAddr returns a local address that was free when checked.
func freeAddr(t *testing.T) string {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	addr := l.Addr().String()
	require.NoError(t, l.Close())
	return addr
}

// serve runs srv until the returned function is called. Serve's result is sent on the channel.
func serve(t *testing.T, srv server.Server) (context.CancelFunc, <-chan error) {
	t.Helper()
	ctx, cancel := context.WithCancel(t.Context())
	serveErr := make(chan error, 1)
	go func() { serveErr <- srv.Serve(ctx) }()
	return cancel, serveErr
}

// waitListening waits until addr accepts TCP connections.
func waitListening(t *testing.T, addr string) {
	t.Helper()
	require.Eventually(t, func() bool {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return false
		}
		_ = conn.Close()
		return true
	}, 5*time.Second, 20*time.Millisecond, "server should listen on %s", addr)
}

// connect opens a Streamable HTTP client session to the server at addr.
// With standaloneSSE, the client also holds a GET stream open for server-initiated messages.
func connect(t *testing.T, addr string, standaloneSSE bool) *mcp.ClientSession {
	t.Helper()
	waitListening(t, addr)
	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "0.0.0"}, nil)
	transport := &mcp.StreamableClientTransport{
		Endpoint:             "http://" + addr + "/mcp",
		MaxRetries:           -1,
		DisableStandaloneSSE: !standaloneSSE,
	}
	session, err := client.Connect(t.Context(), transport, nil)
	require.NoError(t, err, "client should connect over streamable HTTP")
	return session
}

// recv waits for a value from ch, failing the test if none arrives in time.
func recv[T any](t *testing.T, ch <-chan T, what string) T {
	t.Helper()
	select {
	case v := <-ch:
		return v
	case <-time.After(5 * time.Second):
		t.Fatalf("timed out waiting for %s", what)
		panic("unreachable")
	}
}

// funcTool is a tool that runs fn when called.
type funcTool struct {
	name string
	fn   func(ctx context.Context)
}

func (f *funcTool) Tool() *mcp.Tool {
	return &mcp.Tool{Name: f.name}
}

func (f *funcTool) Register(s *mcp.Server) {
	mcp.AddTool(s, f.Tool(), func(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, any, error) {
		if f.fn != nil {
			f.fn(ctx)
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: "done"}}}, nil, nil
	})
}

// TestServeStreamableHTTP_ToolAnnotations checks that tool annotations survive
// registration and serialization, so clients see them in tools/list.
func TestServeStreamableHTTP_ToolAnnotations(t *testing.T) {
	addr := freeAddr(t)
	tools := []server.Tool{tool_query.NewQuery(databasemock.NewPoolMock(), tool_query.WithReadOnly())}
	srv := server.NewServer(slogt.New(t), "0.0.0", "streamable-http", addr, tools, nil, nil)
	cancel, serveErr := serve(t, srv)

	session := connect(t, addr, false)
	res, err := session.ListTools(t.Context(), nil)
	require.NoError(t, err)
	require.Len(t, res.Tools, 1)
	require.NotNil(t, res.Tools[0].Annotations)
	assert.True(t, res.Tools[0].Annotations.ReadOnlyHint)
	require.NoError(t, session.Close())

	cancel()
	require.NoError(t, <-serveErr, "Serve should shut down cleanly")
}

// mockTool is a mock implementation of the Tool interface for testing.
type mockTool struct {
	mock.Mock
}

func (m *mockTool) Tool() *mcp.Tool {
	args := m.Called()
	return args.Get(0).(*mcp.Tool)
}

func (m *mockTool) Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mcp.CallToolResult), args.Error(1)
}

func (m *mockTool) Register(s *mcp.Server) {
	// call Tool
	_ = m.Tool()
}

// mockPrompt is a mock implementation of the Prompt interface for testing.
type mockPrompt struct {
	mock.Mock
}

func (m *mockPrompt) Prompt() *mcp.Prompt {
	args := m.Called()
	return args.Get(0).(*mcp.Prompt)
}

func (m *mockPrompt) Handler(ctx context.Context, request *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
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

func (m *mockResourceTemplate) ResourceTemplate() *mcp.ResourceTemplate {
	args := m.Called()
	return args.Get(0).(*mcp.ResourceTemplate)
}

func (m *mockResourceTemplate) Handler(ctx context.Context, request *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	args := m.Called(ctx, request)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mcp.ReadResourceResult), args.Error(1)
}
