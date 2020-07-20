package oauth2_oidc_sdk

import (
	"net/http"
	"net/url"
	"time"
)

type (
	IBaseContext interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetClientID() string
		SetClient(client IClient)
		GetSecret() string
		GetClient() IClient
		GetError() IError
		SetError(err IError)
		GetForm() *url.Values
	}
	IRevocationRequestContext interface {
		IBaseContext
		GetToken() string
		GetTokenTypeHint() string
	}
	RevocationRequestContextFactory func(r *http.Request) (IRevocationRequestContext, IError)
	RevocationResponseWriter        func(requestContext IRevocationRequestContext, w http.ResponseWriter, r *http.Request) error

	IIntrospectionRequestContext interface {
		IBaseContext
		GetProfile() RequestProfile
		SetProfile(profile RequestProfile)
		GetToken() string
		GetTokenTypeHint() string
	}
	IntrospectionRequestContextFactory func(r *http.Request) (IRevocationRequestContext, IError)
	IntrospectionResponseWriter        func(requestContext IRevocationRequestContext, w http.ResponseWriter, r *http.Request) error
)

//mysqldump --user=root --password=root --result-file=dump.sql --databases db