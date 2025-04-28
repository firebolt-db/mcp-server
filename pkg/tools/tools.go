package tools

import "github.com/mark3labs/mcp-go/mcp"

// textOrResourceContent returns a text content if disableResources is true, otherwise returns an embedded resource.
func textOrResourceContent(disableResources bool, i mcp.ResourceContents) mcp.Content {

	if disableResources {
		textResource, ok := i.(mcp.TextResourceContents)
		if ok {
			return mcp.NewTextContent(textResource.Text)
		}
	}

	return mcp.NewEmbeddedResource(i)
}
