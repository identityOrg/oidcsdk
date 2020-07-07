package oauth2_oidc_sdk

type (
	IRevocationRequestContext interface {
		GetClientID() string
		GetClient() IClient
		GetSecret() IClient
		GetAuthorizationToken() string
		GetToken() string
		GetTokenTypeHint() string
	}
)
