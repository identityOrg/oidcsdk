package oauth2_oidc_sdk

import (
	"time"
)

type (
	Config struct {
		Issuer               string
		AuthCodeLifespan     time.Duration
		AccessTokenLifespan  time.Duration
		RefreshTokenLifespan time.Duration
	}
	IConfigurable interface {
		Configure(strategy interface{}, config Config, arg ...interface{})
	}

	IError interface {
		error
		GetErrorCode() string
		GetDescription() string
		GetStatusCode() int
		GetErrorURL() string
		WithDescription(desc string) IError
	}

	ErrorFactory func(status uint8, code string, description string) IError
)
