package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"

	"github.com/urfave/cli/v3"

	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/prompts"
	"github.com/firebolt-db/mcp-server/pkg/server"
)

var (
	version = "0.0.0"
	commit  = "unknown"
	date    = "unknown"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cmd := &cli.Command{
		Name:    "firebolt-mcp-server",
		Usage:   "Model Context Protocol implementation that connects your LLM to Firebolt",
		Version: fullVersion(),
		Authors: []any{"Firebolt Team"},
		Description: "" +
			"This MCP makes your LLM an expert in Firebolt cloud data warehouse with access to specialized tools and resources.\n" +
			"It can assist with SQL queries, data modeling, performance optimization, and analytics for Firebolt.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "transport",
				Category: "MCP Transport",
				Value:    "stdio",
				Usage:    "Transport type (stdio or sse)",
				Sources:  cli.EnvVars("FIREBOLT_MCP_TRANSPORT"),
				Validator: func(s string) error {
					if s != "stdio" && s != "sse" {
						return fmt.Errorf("invalid transport type: %s, valid options: stdio, sse", s)
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:     "transport-sse-listen-address",
				Category: "MCP Transport",
				Value:    ":8080",
				Usage:    "SSE transport listen address (used only if transport is set to sse)",
				Sources:  cli.EnvVars("FIREBOLT_MCP_TRANSPORT_SSE_LISTEN_ADDRESS"),
			},
			&cli.BoolFlag{
				Name:     "disable-resources",
				Category: "MCP Transport",
				Value:    false,
				Usage:    "Return text content instead of embedded resources (for clients that do not support resources)",
				Sources:  cli.EnvVars("FIREBOLT_MCP_DISABLE_RESOURCES"),
			},
			&cli.StringFlag{
				Name:     "client-id",
				Category: "Firebolt Authentication",
				Value:    "",
				Usage:    "Service account client ID for authentication",
				Sources:  cli.EnvVars("FIREBOLT_MCP_CLIENT_ID"),
			},
			&cli.StringFlag{
				Name:     "client-secret",
				Category: "Firebolt Authentication",
				Value:    "",
				Usage:    "Service account client secret for authentication",
				Sources:  cli.EnvVars("FIREBOLT_MCP_CLIENT_SECRET"),
			},
			&cli.StringFlag{
				Name:     "environment",
				Category: "Firebolt Environment",
				Hidden:   true,
				Value:    "app.firebolt.io",
				Usage:    "Firebolt environment to connect to",
				Sources:  cli.EnvVars("FIREBOLT_MCP_ENVIRONMENT"),
			},
			&cli.StringFlag{
				Name:     "core-url",
				Category: "Firebolt Environment",
				Value:    "",
				Usage:    "Firebolt Core URL to connect to",
				Sources:  cli.EnvVars("FIREBOLT_MCP_CORE_URL"),
			},
			&cli.BoolFlag{
				Name:     "skip-docs-proof",
				Category: "MCP Tools Configuration",
				Value:    false,
				Usage: "Skip the requirement for LLM to provide a token as a proof it has reviewed documentation overview. When enabled, LLM " +
					"will not be forced to gather more starting context and become smarter, but this means more tokens consumed and slower responses.",
				Sources: cli.EnvVars("FIREBOLT_MCP_SKIP_DOCS_PROOF"),
			},
		},
		Action: run,
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context, cmd *cli.Command) error {

	// Initialize logger
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	cfg := server.NewConfig(cmd)
	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Connect to Firebolt
	err := os.Setenv("FIREBOLT_ENDPOINT", fmt.Sprintf("https://api.%s", cfg.Environment))
	if err != nil {
		return fmt.Errorf("failed to set FIREBOLT_ENDPOINT environment variable: %w", err)
	}

	dbPool, dbPoolClose := getDBPool(cfg, logger)
	defer dbPoolClose()

	// prepare tools and resource templates
	tools, resourceTemplates, err := prepareToolsAndResourceTemplates(ctx, logger, cfg, dbPool)
	if err != nil {
		return fmt.Errorf("failed to prepare tools and resource templates: %w", err)
	}

	// Initialize MCP server
	srv := server.NewServer(
		logger,
		fullVersion(),
		cfg.Transport,
		cfg.TransportSSEListenAddress,
		tools,
		[]server.Prompt{
			prompts.NewFireboltExpert(),
		},
		resourceTemplates,
	)

	// Start the server
	logger.Info("Welcome to Firebolt MCP Server!", "version", version)
	if err = srv.Serve(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			logger.Info("Server stopped")
			return nil
		}
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// getDBPool returns a database pool for the given Firebolt environment.
// It also checks provided configuration values.
// Required values for Firebolt:
// - Client ID
// - Client Secret
// Required values for Firebolt Core:
// - Core URL
func getDBPool(cfg server.Config, logger *slog.Logger) (database.Pool, func()) {
	isCore := cfg.CoreURL != ""

	// check if the configuration specifies Firebolt Core connection
	if isCore {
		return database.NewCorePool(logger, cfg.CoreURL)
	}

	return database.NewPool(logger, cfg.ClientID, cfg.ClientSecret)
}

// generateRandomSecret generates a random 32-character alphanumeric string.
func generateRandomSecret() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func fullVersion() string {

	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		return fmt.Sprintf(
			"%s (%s/%s, %s, %s, %s)",
			version,
			runtime.GOOS,
			runtime.GOARCH,
			buildInfo.GoVersion,
			date,
			commit,
		)
	}

	return fmt.Sprintf(
		"%s (%s/%s, %s, %s)",
		version,
		runtime.GOOS,
		runtime.GOARCH,
		date,
		commit,
	)
}
