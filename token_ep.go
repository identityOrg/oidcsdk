package oauth2_oidc_sdk

import (
	"net/http"
	"net/url"
	"time"
)

type (
	ITokenRequestContext interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetState() string
		GetRedirectURI() string
		GetGrantType() string
		GetClientID() string
		GetClientSecret() string
		GetAuthorizationCode() string
		GetRefreshToken() string
		GetUsername() string
		GetPassword() string
		GetRequestedScopes() Arguments
		GetRequestedAudience() Arguments
		GetGrantedScopes() Arguments
		GetGrantedAudience() Arguments
		GrantScope(scope string)
		GrantAudience(audience string)
		GetClient() IClient
		SetClient(client IClient)
		GetProfile() IProfile
		SetProfile(profile IProfile)
		GetIssuedTokens() *Tokens
		GetForm() *url.Values
	}

	TokenRequestContextFactory func(r *http.Request) (ITokenRequestContext, IError)
	TokenResponseWriter        func(requestContext ITokenRequestContext, w http.ResponseWriter) error
	TokenErrorWriter           func(err IError, w http.ResponseWriter) error
)
