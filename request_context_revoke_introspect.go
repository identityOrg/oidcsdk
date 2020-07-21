package oidcsdk

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
		GetToken() string
		GetTokenTypeHint() string
		SetClient(client IClient)
		GetClientSecret() string
		GetClient() IClient
		GetError() IError
		SetError(err IError)
		GetForm() *url.Values
	}
	IRevocationRequestContext interface {
		IBaseContext
	}
	RevocationRequestContextFactory func(request *http.Request) (IRevocationRequestContext, IError)
	RevocationResponseWriter        func(requestContext IRevocationRequestContext, writer http.ResponseWriter, request *http.Request) error

	IIntrospectionRequestContext interface {
		IBaseContext
		GetProfile() RequestProfile
		SetProfile(profile RequestProfile)
		IsActive() bool
		SetActive(active bool)
		GetTokenType() string
		SetTokenType(tokenType string)
	}
	IntrospectionRequestContextFactory func(request *http.Request) (IIntrospectionRequestContext, IError)
	IntrospectionResponseWriter        func(requestContext IIntrospectionRequestContext, writer http.ResponseWriter, request *http.Request) error
	BearerErrorResponseWriter          func(writer http.ResponseWriter, request *http.Request) error
)

//mysqldump --user=root --password=root --result-file=dump.sql --databases db
