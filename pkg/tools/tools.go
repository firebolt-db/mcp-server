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
