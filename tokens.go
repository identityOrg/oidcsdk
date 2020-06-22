package oauth2_oidc_sdk

import (
	"time"
)

type (
	ITokens interface {
		ITokenSignatures
		GetAuthorizationCode() string
		GetAccessToken() string
		GetRefreshToken() string
		GetAccessTokenExpiry() time.Duration
		GetTokenType() string
		GetIDToken() string
	}
	ITokenSignatures interface {
		GetAuthorizationCodeSignature() string
		GetAccessTokenSignature() string
		GetRefreshTokenSignature() string
	}
	IAuthorizationCodeStrategy interface {
		GenerateAuthCode() (code string, signature string)
		ValidateAuthCode(code string, signature string) error
		SignAuthCode(code string) string
	}
	IAccessTokenStrategy interface {
		GenerateAccessToken() (token string, signature string)
		ValidateAccessToken(token string, signature string) error
	}
	IRefreshTokenStrategy interface {
		GenerateRefreshToken() (token string, signature string)
		ValidateRefreshToken(token string, signature string) error
		SignRefreshToken(token string) string
	}
	IIDTokenStrategy interface {
		GenerateIDToken(profile IProfile, client IClient, transactionClaims map[string]interface{}) (idToken string, err error)
	}
	TokensFactory func() ITokens
)
