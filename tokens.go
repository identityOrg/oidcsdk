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
		RefreshTokenExpiry         time.Time
		AccessTokenExpiry          time.Time
		AuthorizationCodeExpiry    time.Time
	}
	IAuthorizationCodeStrategy interface {
		GenerateAuthCode() (code string, signature string)
		SignAuthCode(token string) (signature string, err error)
	}
	IAccessTokenStrategy interface {
		GenerateAccessToken() (token string, signature string)
		SignAccessToken(token string) (signature string, err error)
	}
	IRefreshTokenStrategy interface {
		GenerateRefreshToken() (token string, signature string)
		SignRefreshToken(token string) (signature string, err error)
	}
	IIDTokenStrategy interface {
		GenerateIDToken(profile RequestProfile, client IClient, expiry time.Time,
			transactionClaims map[string]interface{}) (idToken string, err error)
	}
)
