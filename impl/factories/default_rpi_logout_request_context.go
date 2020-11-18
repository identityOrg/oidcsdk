package factories

import (
	sdk "github.com/identityOrg/oidcsdk"
)

type DefaultRPILogoutRequestContext struct {
	RedirectUri string
	State       string
	Client      sdk.IClient
	Username    string
	CSRFToken   string
	Token       string
	Session     sdk.ISession
}

func (d *DefaultRPILogoutRequestContext) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultRPILogoutRequestContext) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultRPILogoutRequestContext) SetUsername(username string) {
	d.Username = username
}

func (d *DefaultRPILogoutRequestContext) GetUserName() string {
	return d.Username
}

func (d *DefaultRPILogoutRequestContext) SetUserSession(session sdk.ISession) {
	d.Session = session
}

func (d *DefaultRPILogoutRequestContext) GetUserSession() sdk.ISession {
	return d.Session
}

func (d *DefaultRPILogoutRequestContext) GetPostLogoutRedirectUri() string {
	return d.RedirectUri
}

func (d *DefaultRPILogoutRequestContext) SetPostLogoutRedirectUri(uri string) {
	d.RedirectUri = uri
}

func (d *DefaultRPILogoutRequestContext) GetIdTokenHint() string {
	return d.Token
}

func (d *DefaultRPILogoutRequestContext) GetState() string {
	return d.State
}

func (d *DefaultRPILogoutRequestContext) GetCSRFToken() string {
	return d.CSRFToken
}
