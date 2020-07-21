package oidcsdk

import (
	"net/http"
	"net/url"
	"time"
)

type (
	IClientCredentialContext interface {
		GetClientID() string
		GetClientSecret() string
		SetClient(client IClient)
	}
	IRequestContext interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetState() string
		GetRedirectURI() string
		GetClientID() string
		GetRequestedScopes() Arguments
		GetRequestedAudience() Arguments
		GetClaims() map[string]interface{}
		GetClient() IClient
		SetClient(client IClient)
		GetProfile() RequestProfile
		SetProfile(profile RequestProfile)
		GetIssuedTokens() Tokens
		IssueAccessToken(token string, signature string, expiry time.Time)
		IssueRefreshToken(token string, signature string, expiry time.Time)
		IssueIDToken(token string)
		GetError() IError
		SetError(err IError)
		GetForm() *url.Values
	}
	IAuthenticationRequestContext interface {
		IRequestContext
		GetUserSession() ISession
		SetUserSession(sess ISession)
		GetNonce() string
		GetResponseMode() string
		GetResponseType() Arguments
		SetRedirectURI(uri string)
		IssueAuthorizationCode(code string, signature string, expiry time.Time)
	}
	ITokenRequestContext interface {
		IRequestContext
		GetRefreshToken() string
		GetPreviousRequestID() (id string)
		SetPreviousRequestID(id string)
		GetGrantType() string
		GetClientSecret() string
		GetAuthorizationCode() string
		GetUsername() string
		GetPassword() string
	}
	TokenRequestContextFactory          func(request *http.Request) (ITokenRequestContext, IError)
	TokenResponseWriter                 func(requestContext ITokenRequestContext, w http.ResponseWriter, r *http.Request) error
	JsonErrorWriter                     func(pError IError, w http.ResponseWriter, r *http.Request) error
	AuthenticationRequestContextFactory func(request *http.Request) (IAuthenticationRequestContext, IError)
	AuthenticationResponseWriter        func(requestContext IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error
	RedirectErrorWriter                 func(requestContext IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error
)
