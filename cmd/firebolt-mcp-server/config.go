package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/urfave/cli/v3"
)

type config struct {
	transport                 string
	transportSSEListenAddress string
	disableResources          bool

	clientID     string
	clientSecret string

	environment string
	coreURL     string

	skipDocsProof bool
}

func newConfig(cmd *cli.Command) config {
	return config{
		transport:                 cmd.String("transport"),
		transportSSEListenAddress: cmd.String("transport-sse-listen-address"),
		disableResources:          cmd.Bool("disable-resources"),
		clientID:                  cmd.String("client-id"),
		clientSecret:              cmd.String("client-secret"),
		environment:               cmd.String("environment"),
		coreURL:                   cmd.String("core-url"),
		skipDocsProof:             cmd.Bool("skip-docs-proof"),
	}
}

func (c config) validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.transport, validation.Required),
		validation.Field(&c.transportSSEListenAddress, validation.Required),
		validation.Field(&c.clientID, validation.When(c.coreURL == "", validation.Required)),
		validation.Field(&c.clientSecret, validation.When(c.coreURL == "", validation.Required)),
		validation.Field(&c.environment, validation.Required),
		validation.Field(&c.coreURL,
			validation.When(c.coreURL != "",
				validation.Required,
				is.URL,
			),
		),
	)
}
