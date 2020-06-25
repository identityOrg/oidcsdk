package tokenep

import (
	"net/url"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/util"
	"time"
)

type (
	DefaultTokenRequestContext struct {
		RequestID         string
		PreviousRequestID string
		RequestedAt       time.Time
		State             string
		RedirectURI       string
		GrantType         string
		ClientId          string
		ClientSecret      string
		Username          string
		Password          string
		AuthorizationCode string
		RefreshToken      string
		RequestedScopes   sdk.Arguments
		RequestedAudience sdk.Arguments
		GrantedScopes     sdk.Arguments
		GrantedAudience   sdk.Arguments
		Client            sdk.IClient
		Profile           sdk.IProfile
		IssuedTokens      sdk.Tokens
		Error             sdk.IError
		Form              *url.Values
	}
)

func (d *DefaultTokenRequestContext) GetError() sdk.IError {
	return d.Error
}

func (d *DefaultTokenRequestContext) SetError(err sdk.IError) {
	d.Error = err
}

func (d *DefaultTokenRequestContext) SetPreviousRequestID(id string) {
	d.PreviousRequestID = id
}

func (d *DefaultTokenRequestContext) GetPreviousRequestID() (id string) {
	return d.PreviousRequestID
}

func (d *DefaultTokenRequestContext) GetIssuedTokens() sdk.Tokens {
	return d.IssuedTokens
}

func (d *DefaultTokenRequestContext) IssueAccessToken(token string, signature string, expiry time.Time) {
	d.IssuedTokens.AccessToken = token
	d.IssuedTokens.AccessTokenSignature = signature
	d.IssuedTokens.AccessTokenExpiry = expiry
}

func (d *DefaultTokenRequestContext) IssueAuthorizationCode(code string, signature string, expiry time.Time) {
	d.IssuedTokens.AuthorizationCode = code
	d.IssuedTokens.AuthorizationCodeSignature = signature
	d.IssuedTokens.AuthorizationCodeExpiry = expiry
}

func (d *DefaultTokenRequestContext) IssueRefreshToken(token string, signature string, expiry time.Time) {
	d.IssuedTokens.RefreshToken = token
	d.IssuedTokens.RefreshTokenSignature = signature
	d.IssuedTokens.RefreshTokenExpiry = expiry
}

func (d *DefaultTokenRequestContext) IssueIDToken(token string) {
	d.IssuedTokens.IDToken = token
}

func (d *DefaultTokenRequestContext) GetUsername() string {
	return d.Username
}

func (d *DefaultTokenRequestContext) GetPassword() string {
	return d.Password
}

func (d *DefaultTokenRequestContext) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultTokenRequestContext) GetGrantedScopes() sdk.Arguments {
	return d.GrantedScopes
}

func (d *DefaultTokenRequestContext) GetGrantedAudience() sdk.Arguments {
	return d.GrantedAudience
}

func (d *DefaultTokenRequestContext) GrantScope(scope string) {
	d.GrantedScopes = util.AppendUnique(d.GrantedScopes, scope)
}

func (d *DefaultTokenRequestContext) GrantAudience(audience string) {
	d.GrantedAudience = util.AppendUnique(d.GrantedAudience, audience)
}

func (d *DefaultTokenRequestContext) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultTokenRequestContext) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultTokenRequestContext) GetProfile() sdk.IProfile {
	return d.Profile
}

func (d *DefaultTokenRequestContext) SetProfile(profile sdk.IProfile) {
	d.Profile = profile
}

func (d *DefaultTokenRequestContext) GetForm() *url.Values {
	return d.Form
}

func (d *DefaultTokenRequestContext) GetRequestedAt() time.Time {
	return d.RequestedAt
}

func (d *DefaultTokenRequestContext) GetState() string {
	return d.State
}

func (d *DefaultTokenRequestContext) GetRedirectURI() string {
	return d.RedirectURI
}

func (d *DefaultTokenRequestContext) GetGrantType() string {
	return d.GrantType
}

func (d *DefaultTokenRequestContext) GetClientID() string {
	return d.ClientId
}

func (d *DefaultTokenRequestContext) GetClientSecret() string {
	return d.ClientSecret
}

func (d *DefaultTokenRequestContext) GetAuthorizationCode() string {
	return d.AuthorizationCode
}

func (d *DefaultTokenRequestContext) GetRefreshToken() string {
	return d.RefreshToken
}

func (d *DefaultTokenRequestContext) GetRequestedScopes() sdk.Arguments {
	return d.RequestedScopes
}

func (d *DefaultTokenRequestContext) GetRequestedAudience() sdk.Arguments {
	return d.RequestedAudience
}
