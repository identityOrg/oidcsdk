package oauth2_oidc_sdk

import (
	"net/http"
	"net/url"
	"time"
)

type (
	IAuthenticationRequestContext interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetState() string
		GetClientID() string
		GetRedirectURI() string
		GetResponseType() Arguments
		GetResponseMode() string
		GetNonce() string
		GetRequestedScopes() Arguments
		GetRequestedAudience() Arguments
		GetForm() *url.Values
		SetRedirectURI(uri string)
		GetGrantedScopes() Arguments
		GetGrantedAudience() Arguments
		GrantScope(scope string)
		GrantAudience(audience string)
		GetClient() IClient
		SetClient(client IClient)
		GetProfile() IProfile
		SetProfile(profile IProfile)
		GetIssuedTokens() *Tokens
	}

	AuthenticationRequestContextFactory func(r *http.Request) (IAuthenticationRequestContext, IError)
	AuthenticationResponseWriter        func(requestContext IAuthenticationRequestContext, w http.ResponseWriter) error
	AuthenticationErrorWriter           func(err IError, w http.ResponseWriter) error
)
