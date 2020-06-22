package oauth2_oidc_sdk

type (
	IProfile interface {
		GetUsername() string
		GetSubject() string
		GetClaims() map[string]interface{}
	}

	ProfileFactory func(username string) IProfile
)
