package prompts

import (
	"context"
	_ "embed"

	"github.com/mark3labs/mcp-go/mcp"
)

//go:embed firebolt.md
var fireboltMD string

const fireboltPromptDescription = "Expert that can provide any assistance with Firebolt cloud data warehouse"

type FireboltExpert struct {
}

func NewFireboltExpert() *FireboltExpert {
	return &FireboltExpert{}
}

func (p *FireboltExpert) Prompt() mcp.Prompt {
	return mcp.NewPrompt(
		"Firebolt Expert",
		mcp.WithPromptDescription(fireboltPromptDescription),
	)
}

func (p *FireboltExpert) Handler(_ context.Context, _ mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {

	var messages []mcp.PromptMessage
	messages = append(messages, mcp.NewPromptMessage(
		mcp.RoleAssistant,
		mcp.NewTextContent(fireboltMD),
	))

	return mcp.NewGetPromptResult(fireboltPromptDescription, messages), nil
}
