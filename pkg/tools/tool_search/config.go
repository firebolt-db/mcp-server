package tool_search

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Config struct {
	TopK                         int
	MinScore                     float64
	IncludeChunks, IncludeScores bool
	BaseURL                      string

	ClientID, ClientSecret string
	TokenURL               string
}

func (c Config) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.TopK, validation.Required, validation.Min(1)),
		validation.Field(&c.BaseURL, validation.Required, is.URL),
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
		validation.Field(&c.TokenURL, validation.Required, is.URL),
	)
}
