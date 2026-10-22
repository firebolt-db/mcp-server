package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// TextOrResourceContent returns a text content if disableResources is true, otherwise returns an embedded resource.
func TextOrResourceContent(disableResources bool, i *mcp.ResourceContents) mcp.Content {
	if disableResources {
		return &mcp.TextContent{
			Text: i.Text,
			Meta: i.Meta,
		}
	}

	return &mcp.EmbeddedResource{
		Resource: i,
	}
}

// ReadOnlyAnnotations marks a tool that does not modify its environment.
func ReadOnlyAnnotations() *mcp.ToolAnnotations {
	return &mcp.ToolAnnotations{ReadOnlyHint: true}
}

// Ptr returns a pointer to v, for optional annotation fields.
func Ptr[T any](v T) *T {
	return &v
}
