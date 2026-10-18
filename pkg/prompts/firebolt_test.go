package prompts_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/firebolt-db/mcp-server/pkg/prompts"
)

// TestFireboltExpertName checks that the prompt name is a valid identifier.
// MCP gateways such as Obot reject prompt names outside [A-Za-z0-9_.-].
func TestFireboltExpertName(t *testing.T) {
	prompt := prompts.NewFireboltExpert().Prompt()

	assert.Regexp(t, regexp.MustCompile(`^[A-Za-z0-9_.-]+$`), prompt.Name)
	assert.Equal(t, "Firebolt Expert", prompt.Title)
}
