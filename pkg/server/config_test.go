package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/firebolt-db/mcp-server/pkg/server"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	cfg := fakeConfig()

	cases := map[string]struct {
		setup   func(*server.Config)
		isValid bool
	}{
		"valid Firebolt core config": {
			setup: func(c *server.Config) {
				c.CoreURL = "https://core.firebolt.io"
				c.ClientID = ""
				c.ClientSecret = ""
			},
			isValid: true,
		},
		"valid Firebolt SAAS config": {
			setup: func(c *server.Config) {
				c.CoreURL = ""
			},
			isValid: true,
		},
		"missing transport": {
			setup: func(c *server.Config) {
				c.Transport = ""
			},
			isValid: false,
		},
		"missing transport SSE listen address": {
			setup: func(c *server.Config) {
				c.Transport = "sse"
				c.TransportSSEListenAddress = ""
			},
		},
		"missing environment": {
			setup: func(c *server.Config) {
				c.Environment = ""
			},
			isValid: false,
		},
		"core url is not a URL": {
			setup: func(c *server.Config) {
				c.CoreURL = "not a URL"
			},
			isValid: false,
		},
		"saas - missing client_id": {
			setup: func(c *server.Config) {
				c.CoreURL = ""
				c.ClientID = ""
			},
			isValid: false,
		},
		"saas - missing client_secret": {
			setup: func(c *server.Config) {
				c.CoreURL = ""
				c.ClientSecret = ""
			},
			isValid: false,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := cfg
			test.setup(&cfg)

			err := cfg.Validate()
			if test.isValid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func fakeConfig() server.Config {
	return server.Config{
		Transport:                 "stdio",
		TransportSSEListenAddress: ":8080",
		ClientID:                  "cl_id",
		ClientSecret:              "cl_secret",
		Environment:               "local.firebolt.io",
		CoreURL:                   "http://localhost:9999",
	}
}
