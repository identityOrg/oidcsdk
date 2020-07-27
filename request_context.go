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

	IRequestContextFactory interface {
		BuildTokenRequestContext(request *http.Request) (ITokenRequestContext, IError)
		BuildAuthorizationRequestContext(request *http.Request) (IAuthenticationRequestContext, IError)
		BuildRevocationRequestContext(request *http.Request) (IRevocationRequestContext, IError)
		BuildIntrospectionRequestContext(request *http.Request) (IIntrospectionRequestContext, IError)
	}

	IErrorWriter interface {
		WriteJsonError(pError IError, additionalValues url.Values, w http.ResponseWriter, r *http.Request) error
		WriteRedirectError(requestContext IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error
		WriteBearerError(pError IError, additionalValues url.Values, w http.ResponseWriter, r *http.Request) error
	}

	IResponseWriter interface {
		WriteTokenResponse(requestContext ITokenRequestContext, w http.ResponseWriter, r *http.Request) error
		WriteAuthorizationResponse(requestContext IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error
		WriteIntrospectionResponse(requestContext IIntrospectionRequestContext, w http.ResponseWriter, r *http.Request) error
		WriteRevocationResponse(w http.ResponseWriter, r *http.Request) error
	}

	IRevocationRequestContext interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetClientID() string
		GetToken() string
		GetTokenTypeHint() string
		SetClient(client IClient)
		GetClientSecret() string
		GetClient() IClient
		GetError() IError
		SetError(err IError)
		GetForm() *url.Values
	}

	IIntrospectionRequestContext interface {
		IRevocationRequestContext
		GetProfile() RequestProfile
		SetProfile(profile RequestProfile)
		IsActive() bool
		SetActive(active bool)
		GetTokenType() string
		SetTokenType(tokenType string)
	}
)
