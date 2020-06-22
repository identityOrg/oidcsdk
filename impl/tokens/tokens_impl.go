package tokens

import (
	sdk "oauth2-oidc-sdk"
	"time"
)

type (
	DefaultToken struct {
		AuthorizationCode          string
		AccessToken                string
		RefreshToken               string
		AccessTokenExpiry          time.Duration
		TokenType                  string
		IDToken                    string
		AuthorizationCodeSignature string
		AccessTokenSignature       string
		RefreshTokenSignature      string
	}
)

func (dt *DefaultToken) GetAuthorizationCodeSignature() string {
	return dt.AuthorizationCodeSignature
}

func (dt *DefaultToken) GetAccessTokenSignature() string {
	return dt.AccessTokenSignature
}

func (dt *DefaultToken) GetRefreshTokenSignature() string {
	return dt.RefreshTokenSignature
}

func (dt *DefaultToken) GetAuthorizationCode() string {
	return dt.AuthorizationCode
}

func (dt *DefaultToken) GetAccessToken() string {
	return dt.AccessToken
}

func (dt *DefaultToken) GetRefreshToken() string {
	return dt.RefreshToken
}

func (dt *DefaultToken) GetAccessTokenExpiry() time.Duration {
	return dt.AccessTokenExpiry
}

func (dt *DefaultToken) GetTokenType() string {
	return dt.TokenType
}

func (dt *DefaultToken) GetIDToken() string {
	return dt.IDToken
}

func DefaultTokensFactory() sdk.ITokens {
	return &DefaultToken{}
}
