package tool_search

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	BaseURL string

	ClientID, ClientSecret string
	TokenURL               string
}

func (c Config) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.BaseURL, validation.Required, is.URL),
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
		validation.Field(&c.TokenURL, validation.Required, is.URL),
	)
}
