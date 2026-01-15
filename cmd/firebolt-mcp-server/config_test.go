package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	cfg := fakeConfig()

	cases := map[string]struct {
		setup   func(*config)
		isValid bool
	}{
		"valid Firebolt core config": {
			setup: func(c *config) {
				c.coreURL = "https://core.firebolt.io"
				c.clientID = ""
				c.clientSecret = ""
			},
			isValid: true,
		},
		"valid Firebolt SAAS config": {
			setup: func(c *config) {
				c.coreURL = ""
			},
			isValid: true,
		},
		"missing transport": {
			setup: func(c *config) {
				c.transport = ""
			},
			isValid: false,
		},
		"missing transport SSE listen address": {
			setup: func(c *config) {
				c.transportSSEListenAddress = ""
			},
		},
		"missing environment": {
			setup: func(c *config) {
				c.environment = ""
			},
			isValid: false,
		},
		"core url is not a URL": {
			setup: func(c *config) {
				c.coreURL = "not a URL"
			},
			isValid: false,
		},
		"saas - missing client_id": {
			setup: func(c *config) {
				c.coreURL = ""
				c.clientID = ""
			},
			isValid: false,
		},
		"saas - missing client_secret": {
			setup: func(c *config) {
				c.coreURL = ""
				c.clientSecret = ""
			},
			isValid: false,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			cfg := cfg
			test.setup(&cfg)

			err := cfg.validate()
			if test.isValid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func fakeConfig() config {
	return config{
		transport:                 "stdio",
		transportSSEListenAddress: ":8080",
		clientID:                  "cl_id",
		clientSecret:              "cl_secret",
		environment:               "local.firebolt.io",
		coreURL:                   "http://localhost:9999",
	}
}
