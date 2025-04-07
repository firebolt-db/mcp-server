package resources

import (
	"context"
	_ "embed"
	"fmt"
	"io/fs"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/firebolt-db/mcp-server/pkg/helpers/args"
	"github.com/firebolt-db/mcp-server/pkg/helpers/mimetype"
)

// Special articles used by this MCP server.
const (
	DocsArticleOverview  = "mcp/overview.md"
	DocsArticleProof     = "mcp/proof.md"
	DocsArticleReference = "mcp/reference.md"
)

//go:embed docs_overview.md
var docsOverviewMD string

//go:embed docs_proof.md
var docsProofMD string

//go:embed docs_reference.md
var docsReferenceMD string

// DocsURI creates a formatted Firebolt documentation URI for a given article.
func DocsURI(article string) string {
	return fmt.Sprintf("firebolt://docs/%s", article)
}

// DocsFS is an interface that combines the ReadDirFS and ReadFileFS interfaces.
type DocsFS interface {
	fs.ReadDirFS
	fs.ReadFileFS
}

// Docs is a resource template handler for serving Firebolt documentation articles.
type Docs struct {
	docsFS      DocsFS
	proofSecret string
}

// NewDocs creates and returns a new instance of the Docs.
func NewDocs(docsFS DocsFS, proofSecret string) *Docs {
	return &Docs{
		docsFS:      docsFS,
		proofSecret: proofSecret,
	}
}

// ResourceTemplate defines the template for documentation resources.
// It specifies the URI format, content type, description, and suggested usage.
func (r *Docs) ResourceTemplate() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		DocsURI("{article}"),
		"Documentation article",
		mcp.WithTemplateMIMEType(mimetype.Markdown),
		mcp.WithTemplateDescription("An article that offers insights into a particular topic related to Firebolt."),
		mcp.WithTemplateAnnotations([]mcp.Role{mcp.RoleUser, mcp.RoleAssistant}, 0.5),
	)
}

// Handler processes resource requests for documentation articles.
// It extracts the article parameter and fetches the appropriate documentation content.
func (r *Docs) Handler(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {

	// Extract the article name from the request parameters
	value, err := args.String(request.Params.Arguments, "article")
	if err != nil {
		return nil, fmt.Errorf("bad request: %w", err)
	}

	return r.FetchDocsResources(ctx, value)
}

// FetchDocsResources retrieves the content for a specified documentation article.
// It handles special articles (overview, reference) and regular documentation files.
func (r *Docs) FetchDocsResources(_ context.Context, article string) ([]mcp.ResourceContents, error) {

	switch article {

	case DocsArticleOverview:

		// Return the pre-embedded overview markdown content
		return r.newResource(article, docsOverviewMD)

	case DocsArticleProof:

		// Return the pre-embedded proof markdown content
		return r.newResource(article, fmt.Sprintf(docsProofMD, r.proofSecret))

	case DocsArticleReference:

		// Retrieve a list of available docs files
		entries, err := r.docsFS.ReadDir(".")
		if err != nil {
			return nil, fmt.Errorf("failed to read docs directory: %w", err)
		}

		// Build a Markdown list of available files
		b := &strings.Builder{}
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
				continue
			}
			b.WriteString("- " + entry.Name() + "\n")
		}

		// Combine the reference template with the generated file list
		return r.newResource(article, docsReferenceMD+b.String())

	default:

		// For other articles, read the file directly from the filesystem
		fileData, err := r.docsFS.ReadFile(article)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", article, err)
		}

		return r.newResource(article, string(fileData))
	}
}

// newResource creates a properly formatted resource contents object for a documentation article.
func (r *Docs) newResource(article, content string) ([]mcp.ResourceContents, error) {
	return []mcp.ResourceContents{&mcp.TextResourceContents{
		URI:      DocsURI(article),
		MIMEType: mimetype.Markdown,
		Text:     content,
	}}, nil
}
