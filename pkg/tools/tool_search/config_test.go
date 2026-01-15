package tool_search_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/tools/tool_search"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	cfg := newConfig()
	require.NoError(t, cfg.Validate())
}

func TestConfig_Validate_Errors(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup  func(c *tool_search.Config)
		expErr string
	}{
		"missing BaseURL": {
			setup: func(c *tool_search.Config) {
				c.BaseURL = ""
			},
			expErr: "BaseURL: cannot be blank",
		},
		"BaseURL is not a URL": {
			setup: func(c *tool_search.Config) {
				c.BaseURL = "not a URL"
			},
			expErr: "BaseURL: must be a valid URL",
		},
		"missing ClientID": {
			setup: func(c *tool_search.Config) {
				c.ClientID = ""
			},
			expErr: "ClientID: cannot be blank",
		},
		"missing ClientSecret": {
			setup: func(c *tool_search.Config) {
				c.ClientSecret = ""
			},
			expErr: "ClientSecret: cannot be blank",
		},
		"missing TokenURL": {
			setup: func(c *tool_search.Config) {
				c.TokenURL = ""
			},
			expErr: "TokenURL: cannot be blank",
		},
		"TokenURL is not a URL": {
			setup: func(c *tool_search.Config) {
				c.TokenURL = "not a URL"
			},
			expErr: "TokenURL: must be a valid URL",
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := newConfig()
			test.setup(&cfg)
			err := cfg.Validate()
			require.Error(t, err)
			require.ErrorContains(t, err, test.expErr)
		})
	}
}
