package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

// textOrResourceContent returns a text content if disableResources is true, otherwise returns an embedded resource.
func textOrResourceContent(disableResources bool, i mcp.ResourceContents) mcp.Content {

	if disableResources {
		switch resource := i.(type) {
		case mcp.TextResourceContents:
			return mcp.NewTextContent(resource.Text)
		case *mcp.TextResourceContents:
			return mcp.NewTextContent(resource.Text)
		}
	}

	return mcp.NewEmbeddedResource(i)
}
