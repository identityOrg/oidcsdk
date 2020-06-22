package oauth2_oidc_sdk

import (
	"net/http"
	"net/url"
	"time"
)

type (
	IAuthenticationRequest interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetState() string
		GetClientID() string
		GetRedirectURI() string
		GetResponseType() Arguments
		GetNonce() string
		GetRequestedScopes() Arguments
		GetRequestedAudience() Arguments
		GetForm() *url.Values
	}

	IAuthenticationResponse interface {
		ITokenResponse
		GetState() string
		GetRedirectURI() string
		GetResponseMode() string
	}

	AuthenticationRequestFactory  func(r *http.Request) (IAuthenticationRequest, IError)
	AuthenticationResponseFactory func(request IAuthenticationRequest) (IAuthenticationResponse, IError)
	AuthenticationResponseWriter  func(response IAuthenticationResponse, w http.ResponseWriter) error
	AuthenticationErrorWriter     func(err IError, w http.ResponseWriter) error
)
