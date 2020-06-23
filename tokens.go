package oauth2_oidc_sdk

import "time"

type (
	Tokens struct {
		TokenSignatures
		AuthorizationCode string
		AccessToken       string
		RefreshToken      string
		TokenType         string
		IDToken           string
	}
	TokenSignatures struct {
		AuthorizationCodeSignature string
		AccessTokenSignature       string
		RefreshTokenSignature      string
		RefreshTokenExpiry         *time.Time
		AccessTokenExpiry          *time.Time
		AuthorizationCodeExpiry    *time.Time
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
)
