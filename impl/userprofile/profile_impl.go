package userprofile

import (
	sdk "oauth2-oidc-sdk"
)

type DefaultProfile struct {
	Username    string
	RedirectURI string
	Scope       sdk.Arguments
	Audience    sdk.Arguments
}

func (d *DefaultProfile) GetUsername() string {
	return d.Username
}

func (d *DefaultProfile) SetUsername(username string) {
	d.Username = username
}

func (d *DefaultProfile) GetRedirectURI() string {
	return d.RedirectURI
}

func (d *DefaultProfile) SetRedirectURI(uri string) {
	d.RedirectURI = uri
}

func (d *DefaultProfile) GetScope() sdk.Arguments {
	return d.Scope
}

func (d *DefaultProfile) SetScope(scopes sdk.Arguments) {
	d.Scope = scopes
}

func (d *DefaultProfile) GetAudience() sdk.Arguments {
	return d.Audience
}

func (d *DefaultProfile) SetAudience(aud sdk.Arguments) {
	d.Audience = aud
}
