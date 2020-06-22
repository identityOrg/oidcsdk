package oauth2_oidc_sdk

import (
	"net/http"
	"net/url"
	"time"
)

type (
	ITokenRequest interface {
		GetRequestID() string
		GetRequestedAt() time.Time
		GetState() string
		GetRedirectURI() string
		GetGrantType() string
		GetClientId() string
		GetClientSecret() string
		GetAuthorizationCode() string
		GetRefreshToken() string
		GetUsername() string
		GetPassword() string
		GetRequestedScopes() Arguments
		GetRequestedAudience() Arguments
		GetForm() *url.Values
	}

	ITokenResponse interface {
		GetRequestID() string
		GetGrantedScopes() Arguments
		GetGrantedAudience() Arguments
		GrantScope(scope string)
		GrantAudience(audience string)
		IsSuccess() bool
		SetSuccess(success bool)
		GetClient() IClient
		SetClient(client IClient)
		GetProfile() IProfile
		SetProfile(profile IProfile)
		GetIssuedTokens() ITokens
		IssueTokens(tokens ITokens)
		GetError() IError
		SetError(err IError)
		GetForm() *url.Values
	}

	TokenRequestFactory  func(r *http.Request) (ITokenRequest, IError)
	TokenResponseFactory func(request ITokenRequest) (ITokenResponse, IError)
	TokenResponseWriter  func(response ITokenResponse, w http.ResponseWriter) error
	TokenErrorWriter     func(err IError, w http.ResponseWriter) error
)
