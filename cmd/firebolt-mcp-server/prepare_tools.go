package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/firebolt-db/mcp-server/cmd/docs-scrapper/fireboltdocs"
	"github.com/firebolt-db/mcp-server/pkg/clients/database"
	"github.com/firebolt-db/mcp-server/pkg/clients/discovery"
	"github.com/firebolt-db/mcp-server/pkg/resources"
	"github.com/firebolt-db/mcp-server/pkg/server"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_connect"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_connect_core"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_docs"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_query"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_search"
)

// prepareTools prepares the tools for the server based on the provided configuration.
// The set of the tools and their set up differs for Firebolt and Firebolt Core.
func prepareToolsAndResourceTemplates(ctx context.Context, logger *slog.Logger, cfg config, dbPool database.Pool) ([]server.Tool, []server.ResourceTemplate, error) {
	docsProofToken := generateRandomSecret()

	var docsProof *string
	if !cfg.skipDocsProof {
		docsProof = &docsProofToken
	}

	resourceDocs := resources.NewDocs(fireboltdocs.FS, docsProofToken)
	resourceDatabases := resources.NewDatabases(dbPool)

	// Firebolt Core tools
	if cfg.coreURL != "" {
		tools := []server.Tool{
			tool_connect_core.NewConnectCore(resourceDatabases, docsProof, cfg.disableResources),
			tool_docs.NewDocs(resourceDocs, cfg.disableResources),
			tool_query.NewQuery(dbPool),
		}
		resourceTemplates := []server.ResourceTemplate{
			resourceDocs,
			resourceDatabases,
		}

		return tools, resourceTemplates, nil
	}

	discoveryClient, err := discovery.NewClient(
		ctx, logger,
		cfg.clientID, cfg.clientSecret,
		fmt.Sprintf("https://id.%s", cfg.environment),
		fmt.Sprintf("https://api.%s/web/v3", cfg.environment),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Firebolt discovery client: %w", err)
	}

	resourceAccounts := resources.NewAccounts(discoveryClient)
	resourceEngines := resources.NewEngines(dbPool)

	searchCfg := tool_search.Config{
		BaseURL:      fmt.Sprintf("https://api.%s", cfg.environment),
		ClientID:     cfg.clientID,
		ClientSecret: cfg.clientSecret,
		TokenURL:     fmt.Sprintf("https://id.%s/oauth/token", cfg.environment),
	}

	docsSearchTool, err := tool_search.NewSearch(ctx, searchCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create search tool: %w", err)
	}

	// Standard Firebolt tools
	tools := []server.Tool{
		tool_connect.NewConnect(resourceAccounts, resourceDatabases, resourceEngines, docsProof, cfg.disableResources),
		tool_docs.NewDocs(resourceDocs, cfg.disableResources),
		tool_query.NewQuery(dbPool),
		docsSearchTool,
	}

	resourceTemplates := []server.ResourceTemplate{
		resourceDocs,
		resourceAccounts,
		resourceDatabases,
		resourceEngines,
	}

	return tools, resourceTemplates, nil
}
