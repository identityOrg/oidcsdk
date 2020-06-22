package oauth2_oidc_sdk

type (
	IProfile interface {
		GetUsername() string
		GetSubject() string
		GetRedirectURI() string
		GetScope() Arguments
		GetAudience() Arguments
		GetClaims() map[string]interface{}
	}

	ProfileFactory func(username string) IProfile
)
