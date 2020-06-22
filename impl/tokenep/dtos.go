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
		Form              *url.Values
	}
	DefaultTokenResponse struct {
		RequestID       string
		GrantedScopes   sdk.Arguments
		GrantedAudience sdk.Arguments
		Success         bool
		Client          sdk.IClient
		Profile         sdk.IProfile
		IssuedTokens    sdk.ITokens
		Error           sdk.IError
		Form            *url.Values
	}
)

func (d *DefaultTokenRequest) GetUsername() string {
	return d.Username
}

func (d *DefaultTokenRequest) GetPassword() string {
	return d.Password
}

func (d *DefaultTokenResponse) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultTokenResponse) GetGrantedScopes() sdk.Arguments {
	return d.GrantedScopes
}

func (d *DefaultTokenResponse) GetGrantedAudience() sdk.Arguments {
	return d.GrantedAudience
}

func (d *DefaultTokenResponse) GrantScope(scope string) {
	d.GrantedScopes = util.AppendUnique(d.GrantedScopes, scope)
}

func (d *DefaultTokenResponse) GrantAudience(audience string) {
	d.GrantedAudience = util.AppendUnique(d.GrantedAudience, audience)
}

func (d *DefaultTokenResponse) IsSuccess() bool {
	return d.Success
}

func (d *DefaultTokenResponse) SetSuccess(success bool) {
	d.Success = success
}

func (d *DefaultTokenResponse) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultTokenResponse) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultTokenResponse) GetProfile() sdk.IProfile {
	return d.Profile
}

func (d *DefaultTokenResponse) SetProfile(profile sdk.IProfile) {
	d.Profile = profile
}

func (d *DefaultTokenResponse) GetIssuedTokens() sdk.ITokens {
	return d.IssuedTokens
}

func (d *DefaultTokenResponse) IssueTokens(tokens sdk.ITokens) {
	d.IssuedTokens = tokens
}

func (d *DefaultTokenResponse) GetError() sdk.IError {
	return d.Error
}

func (d *DefaultTokenResponse) SetError(err sdk.IError) {
	d.Error = err
}

func (d *DefaultTokenResponse) GetForm() *url.Values {
	return d.Form
}

// starting request

func (d *DefaultTokenRequest) GetRequestID() string {
	return d.RequestID
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

func (d *DefaultTokenRequest) GetForm() *url.Values {
	return d.Form
}
