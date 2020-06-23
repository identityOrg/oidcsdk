package tokenep

import (
	"net/url"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/util"
	"time"
)

type (
	DefaultTokenRequest struct {
		RequestID         string
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
		IssuedTokens      *sdk.Tokens
		Error             sdk.IError
		Form              *url.Values
	}
)

func (d *DefaultTokenRequest) GetUsername() string {
	return d.Username
}

func (d *DefaultTokenRequest) GetPassword() string {
	return d.Password
}

func (d *DefaultTokenRequest) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultTokenRequest) GetGrantedScopes() sdk.Arguments {
	return d.GrantedScopes
}

func (d *DefaultTokenRequest) GetGrantedAudience() sdk.Arguments {
	return d.GrantedAudience
}

func (d *DefaultTokenRequest) GrantScope(scope string) {
	d.GrantedScopes = util.AppendUnique(d.GrantedScopes, scope)
}

func (d *DefaultTokenRequest) GrantAudience(audience string) {
	d.GrantedAudience = util.AppendUnique(d.GrantedAudience, audience)
}

func (d *DefaultTokenRequest) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultTokenRequest) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultTokenRequest) GetProfile() sdk.IProfile {
	return d.Profile
}

func (d *DefaultTokenRequest) SetProfile(profile sdk.IProfile) {
	d.Profile = profile
}

func (d *DefaultTokenRequest) GetIssuedTokens() *sdk.Tokens {
	return d.IssuedTokens
}

func (d *DefaultTokenRequest) GetForm() *url.Values {
	return d.Form
}

func (d *DefaultTokenRequest) GetRequestedAt() time.Time {
	return d.RequestedAt
}

func (d *DefaultTokenRequest) GetState() string {
	return d.State
}

func (d *DefaultTokenRequest) GetRedirectURI() string {
	return d.RedirectURI
}

func (d *DefaultTokenRequest) GetGrantType() string {
	return d.GrantType
}

func (d *DefaultTokenRequest) GetClientID() string {
	return d.ClientId
}

func (d *DefaultTokenRequest) GetClientSecret() string {
	return d.ClientSecret
}

func (d *DefaultTokenRequest) GetAuthorizationCode() string {
	return d.AuthorizationCode
}

func (d *DefaultTokenRequest) GetRefreshToken() string {
	return d.RefreshToken
}

func (d *DefaultTokenRequest) GetRequestedScopes() sdk.Arguments {
	return d.RequestedScopes
}

func (d *DefaultTokenRequest) GetRequestedAudience() sdk.Arguments {
	return d.RequestedAudience
}
