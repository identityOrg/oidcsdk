package factories

import (
	sdk "github.com/identityOrg/oidcsdk"
	"net/url"
	"time"
)

type (
	DefaultAuthenticationRequestContext struct {
		RequestID         string
		RequestedAt       time.Time
		State             string
		RedirectURI       string
		ClientId          string
		Nonce             string
		ResponseMode      string
		RequestedScopes   sdk.Arguments
		RequestedAudience sdk.Arguments
		Claims            map[string]interface{}
		Client            sdk.IClient
		Profile           sdk.RequestProfile
		IssuedTokens      sdk.Tokens
		Error             sdk.IError
		Form              *url.Values
		ResponseType      sdk.Arguments
		UserSession       sdk.ISession
	}
)

func (d *DefaultAuthenticationRequestContext) GetUserSession() sdk.ISession {
	return d.UserSession
}

func (d *DefaultAuthenticationRequestContext) SetUserSession(sess sdk.ISession) {
	d.UserSession = sess
}

func (d *DefaultAuthenticationRequestContext) GetError() sdk.IError {
	return d.Error
}

func (d *DefaultAuthenticationRequestContext) SetError(err sdk.IError) {
	d.Error = err
}

func (d *DefaultAuthenticationRequestContext) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultAuthenticationRequestContext) GetRequestedAt() time.Time {
	return d.RequestedAt
}

func (d *DefaultAuthenticationRequestContext) GetState() string {
	return d.State
}

func (d *DefaultAuthenticationRequestContext) GetRedirectURI() string {
	return d.RedirectURI
}

func (d *DefaultAuthenticationRequestContext) GetClientID() string {
	return d.ClientId
}

func (d *DefaultAuthenticationRequestContext) GetRequestedScopes() sdk.Arguments {
	return d.RequestedScopes
}

func (d *DefaultAuthenticationRequestContext) GetRequestedAudience() sdk.Arguments {
	return d.RequestedAudience
}

func (d *DefaultAuthenticationRequestContext) GetClaims() map[string]interface{} {
	return d.Claims
}

func (d *DefaultAuthenticationRequestContext) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultAuthenticationRequestContext) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultAuthenticationRequestContext) GetProfile() sdk.RequestProfile {
	return d.Profile
}

func (d *DefaultAuthenticationRequestContext) SetProfile(profile sdk.RequestProfile) {
	d.Profile = profile
}

func (d *DefaultAuthenticationRequestContext) GetIssuedTokens() sdk.Tokens {
	return d.IssuedTokens
}

func (d *DefaultAuthenticationRequestContext) IssueAccessToken(token string, signature string, expiry time.Time) {
	d.IssuedTokens.AccessToken = token
	d.IssuedTokens.AccessTokenSignature = signature
	d.IssuedTokens.AccessTokenExpiry = expiry
}

func (d *DefaultAuthenticationRequestContext) IssueAuthorizationCode(code string, signature string, expiry time.Time) {
	d.IssuedTokens.AuthorizationCode = code
	d.IssuedTokens.AuthorizationCodeSignature = signature
	d.IssuedTokens.AuthorizationCodeExpiry = expiry
}

func (d *DefaultAuthenticationRequestContext) IssueRefreshToken(token string, signature string, expiry time.Time) {
	d.IssuedTokens.RefreshToken = token
	d.IssuedTokens.RefreshTokenSignature = signature
	d.IssuedTokens.RefreshTokenExpiry = expiry
}

func (d *DefaultAuthenticationRequestContext) IssueIDToken(token string) {
	d.IssuedTokens.IDToken = token
}

func (d *DefaultAuthenticationRequestContext) GetForm() *url.Values {
	return d.Form
}

func (d *DefaultAuthenticationRequestContext) GetNonce() string {
	return d.Nonce
}

func (d *DefaultAuthenticationRequestContext) GetResponseMode() string {
	if d.ResponseMode != "" {
		return d.ResponseMode
	} else {
		if d.ResponseType.HasOneOf(sdk.ResponseTypeToken, sdk.ResponseTypeIdToken) {
			return sdk.ResponseModeFragment
		} else {
			return sdk.ResponseModeQuery
		}
	}
}

func (d *DefaultAuthenticationRequestContext) GetResponseType() sdk.Arguments {
	return d.ResponseType
}

func (d *DefaultAuthenticationRequestContext) SetRedirectURI(uri string) {
	d.RedirectURI = uri
}
