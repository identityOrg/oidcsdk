package oauth2_oidc_sdk

import (
	"time"
)

type (
	Config struct {
		Issuer                   string
		AuthCodeLifespan         time.Duration
		AccessTokenLifespan      time.Duration
		RefreshTokenLifespan     time.Duration
		AccessTokenEntropy       int
		AuthorizationCodeEntropy int
		RefreshTokenEntropy      int
		StateParamMinimumEntropy int
	}
	IConfigurable interface {
		Configure(strategy interface{}, config *Config, arg ...interface{})
	}

	IError interface {
		error
		GetStatus() string
		GetReason() string
		GetStatusCode() int
		GetDescription() string
		GetDebugInfo() string
	}

	ErrorFactory func(status uint8, code string, description string) IError
)

func NewConfig(issuer string) *Config {
	config := &Config{Issuer: issuer}
	config.StateParamMinimumEntropy = 20
	config.RefreshTokenEntropy = 20
	config.AccessTokenEntropy = 20
	config.AuthorizationCodeEntropy = 20
	config.AuthCodeLifespan = time.Minute * 10
	config.AccessTokenLifespan = time.Minute * 60
	config.RefreshTokenLifespan = time.Hour * 24 * 30
	return config
}
