package prompts

import (
	"context"
	_ "embed"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

//go:embed firebolt.md
var fireboltMD string

const fireboltPromptDescription = "Expert that can provide any assistance with Firebolt cloud data warehouse"

type FireboltExpert struct {
}

func NewFireboltExpert() *FireboltExpert {
	return &FireboltExpert{}
}

func (p *FireboltExpert) Prompt() *mcp.Prompt {
	return &mcp.Prompt{
		Name:        "Firebolt Expert",
		Description: fireboltPromptDescription,
	}
}

func (p *FireboltExpert) Handler(_ context.Context, _ *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {

	var messages []*mcp.PromptMessage
	messages = append(messages, &mcp.PromptMessage{
		Role: "assistant",
		Content: &mcp.TextContent{
			Text: fireboltMD,
		},
	})

	return &mcp.GetPromptResult{
		Description: fireboltPromptDescription,
		Messages:    messages,
	}, nil
}
