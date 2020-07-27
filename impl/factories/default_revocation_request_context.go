package factories

import (
	sdk "github.com/identityOrg/oidcsdk"
	"net/url"
	"time"
)

type DefaultRevocationRequestContext struct {
	RequestID     string
	RequestedAt   time.Time
	ClientID      string
	Token         string
	TokenTypeHint string
	ClientSecret  string
	Client        sdk.IClient
	Error         sdk.IError
	Form          *url.Values
}

func (d *DefaultRevocationRequestContext) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultRevocationRequestContext) GetRequestedAt() time.Time {
	return d.RequestedAt
}

func (d *DefaultRevocationRequestContext) GetClientID() string {
	return d.ClientID
}

func (d *DefaultRevocationRequestContext) GetToken() string {
	return d.Token
}

func (d *DefaultRevocationRequestContext) GetTokenTypeHint() string {
	return d.TokenTypeHint
}

func (d *DefaultRevocationRequestContext) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultRevocationRequestContext) GetClientSecret() string {
	return d.ClientSecret
}

func (d *DefaultRevocationRequestContext) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultRevocationRequestContext) GetError() sdk.IError {
	return d.Error
}

func (d *DefaultRevocationRequestContext) SetError(err sdk.IError) {
	d.Error = err
}

func (d *DefaultRevocationRequestContext) GetForm() *url.Values {
	return d.Form
}
