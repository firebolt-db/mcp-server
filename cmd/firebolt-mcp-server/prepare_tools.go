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
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_docs_llm"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_query"
	"github.com/firebolt-db/mcp-server/pkg/tools/tool_search"
)

// prepareTools prepares the tools for the server based on the provided configuration.
// The set of the tools and their setup differ for Firebolt and Firebolt Core.
func prepareToolsAndResourceTemplates(ctx context.Context, logger *slog.Logger, cfg server.Config, dbPool database.Pool) ([]server.Tool, []server.ResourceTemplate, error) {
	docsProofToken := generateRandomSecret()

	var docsProof *string
	if !cfg.SkipDocsProof {
		docsProof = &docsProofToken
	}

	// resources shared between Firebolt and Firebolt Core
	resourceDocs := resources.NewDocs(fireboltdocs.FS, docsProofToken)

	// Firebolt Core tools and resource templates
	if cfg.CoreURL != "" {
		resourceLLMDocs := resources.NewLLMDocs()
		resourceDatabases := resources.NewDatabases(dbPool, resources.WithCore())
		resourceEngines := resources.NewEngines(dbPool, resources.WithCore())
		resourceCoreAccounts := resources.NewCoreAccounts(dbPool)

		tools := []server.Tool{
			tool_connect_core.NewConnectCore(resourceCoreAccounts, resourceDatabases, resourceEngines, docsProof, cfg.DisableResources),
			tool_docs.NewDocs(resourceDocs, cfg.DisableResources, true),
			tool_docs_llm.NewDocsLLM(resourceLLMDocs, cfg.DisableResources),
			tool_query.NewQuery(dbPool),
		}
		resourceTemplates := []server.ResourceTemplate{
			resourceDocs,
			resourceCoreAccounts,
			resourceDatabases,
			resourceEngines,
			resourceLLMDocs,
		}

		return tools, resourceTemplates, nil
	}

	discoveryClient, err := discovery.NewClient(
		ctx, logger,
		cfg.ClientID, cfg.ClientSecret,
		fmt.Sprintf("https://id.%s", cfg.Environment),
		fmt.Sprintf("https://api.%s/web/v3", cfg.Environment),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Firebolt discovery client: %w", err)
	}

	resourceAccounts := resources.NewAccounts(discoveryClient)
	resourceDatabases := resources.NewDatabases(dbPool)
	resourceEngines := resources.NewEngines(dbPool)

	// Firebolt SaaS tools and resource templates
	searchCfg := tool_search.Config{
		BaseURL:      fmt.Sprintf("https://api.%s", cfg.Environment),
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		TokenURL:     fmt.Sprintf("https://id.%s/oauth/token", cfg.Environment),
	}

	docsSearchTool, err := tool_search.NewSearch(ctx, searchCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create search tool: %w", err)
	}

	// Standard Firebolt tools
	tools := []server.Tool{
		tool_connect.NewConnect(resourceAccounts, resourceDatabases, resourceEngines, docsProof, cfg.DisableResources),
		tool_docs.NewDocs(resourceDocs, cfg.DisableResources, false),
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
