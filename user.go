package oauth2_oidc_sdk

type (
	IProfile interface {
		GetUsername() string
		SetUsername(username string)
		GetRedirectURI() string
		SetRedirectURI(uri string)
		GetScope() Arguments
		SetScope(scopes Arguments)
		GetAudience() Arguments
		SetAudience(aud Arguments)
	}
)
