package server

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
