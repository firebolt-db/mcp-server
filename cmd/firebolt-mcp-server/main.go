package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v3"

	"github.com/firebolt-db/mcp-server/cmd/docs-scrapper/fireboltdocs"
	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/discovery"
	"github.com/firebolt-db/mcp-server/pkg/prompts"
	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/server"
	"github.com/firebolt-db/mcp-server/pkg/tools"
	"github.com/firebolt-db/mcp-server/pkg/version"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cmd := &cli.Command{
		Name:    "firebolt-mcp-server",
		Usage:   "Model Context Protocol implementation that connects your LLM to Firebolt",
		Version: version.GetFullVersion(),
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
			&cli.StringFlag{
				Name:     "client-id",
				Category: "Firebolt Authentication",
				Required: true,
				Value:    "",
				Usage:    "Service account client ID for authentication",
				Sources:  cli.EnvVars("FIREBOLT_MCP_CLIENT_ID"),
			},
			&cli.StringFlag{
				Name:     "client-secret",
				Category: "Firebolt Authentication",
				Required: true,
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

	// Connect to Firebolt
	err := os.Setenv("FIREBOLT_ENDPOINT", fmt.Sprintf("https://api.%s", cmd.String("environment")))
	if err != nil {
		return fmt.Errorf("failed to set FIREBOLT_ENDPOINT environment variable: %w", err)
	}
	clientID := cmd.String("client-id")
	clientSecret := cmd.String("client-secret")
	dbPool, dbPoolClose := database.NewPool(logger, clientID, clientSecret)
	defer dbPoolClose()
	discoveryClient, err := discovery.NewClient(
		ctx, logger,
		clientID, clientSecret,
		fmt.Sprintf("https://id.%s", cmd.String("environment")),
		fmt.Sprintf("https://api.%s/web/v3", cmd.String("environment")),
	)
	if err != nil {
		return fmt.Errorf("failed to create Firebolt discovery client: %w", err)
	}

	// Initialize MCP server
	docsProof := generateRandomSecret()
	resourceDocs := resources.NewDocs(fireboltdocs.FS, docsProof)
	resourceAccounts := resources.NewAccounts(discoveryClient)
	resourceDatabases := resources.NewDatabases(dbPool)
	resourceEngines := resources.NewEngines(dbPool)
	srv := server.NewServer(
		logger,
		cmd.String("transport"),
		cmd.String("transport-sse-listen-address"),
		[]server.Tool{
			tools.NewConnect(resourceAccounts, resourceDatabases, resourceEngines, docsProof),
			tools.NewDocs(resourceDocs),
			tools.NewQuery(dbPool),
		},
		[]server.Prompt{
			prompts.NewFireboltExpert(),
		},
		[]server.ResourceTemplate{
			resourceDocs,
			resourceAccounts,
			resourceDatabases,
			resourceEngines,
		},
	)

	// Start the server
	logger.Info("Welcome to Firebolt MCP Server!", "version", version.GetFullVersion())
	if err = srv.Serve(ctx); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
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
