package server

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/urfave/cli/v3"
)

type Config struct {
	Transport                 string
	TransportSSEListenAddress string
	DisableResources          bool

	ClientID     string
	ClientSecret string

	Environment string
	CoreURL     string

	SkipDocsProof bool
}

func NewConfig(cmd *cli.Command) Config {
	return Config{
		Transport:                 cmd.String("transport"),
		TransportSSEListenAddress: cmd.String("transport-sse-listen-address"),
		DisableResources:          cmd.Bool("disable-resources"),
		ClientID:                  cmd.String("client-id"),
		ClientSecret:              cmd.String("client-secret"),
		Environment:               cmd.String("environment"),
		CoreURL:                   cmd.String("core-url"),
		SkipDocsProof:             cmd.Bool("skip-docs-proof"),
	}
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Transport, validation.Required),
		validation.Field(&c.TransportSSEListenAddress, validation.When(c.Transport == "sse", validation.Required)),
		validation.Field(&c.ClientID, validation.When(c.CoreURL == "", validation.Required)),
		validation.Field(&c.ClientSecret, validation.When(c.CoreURL == "", validation.Required)),
		validation.Field(&c.Environment, validation.Required),
		validation.Field(&c.CoreURL,
			validation.When(c.CoreURL != "",
				validation.Required,
				is.URL,
			),
		),
	)
}
