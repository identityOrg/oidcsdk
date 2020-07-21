package introrevoke

import (
	sdk "github.com/identityOrg/oidcsdk"
	"net/url"
	"time"
)

type DefaultIntrospectionRequestContext struct {
	RequestID     string
	RequestedAt   time.Time
	ClientID      string
	Secret        string
	Client        sdk.IClient
	Error         sdk.IError
	Form          *url.Values
	Profile       sdk.RequestProfile
	Token         string
	TokenTypeHint string
	Active        bool
	TokenType     string
}

func (d *DefaultIntrospectionRequestContext) GetTokenType() string {
	return d.TokenType
}

func (d *DefaultIntrospectionRequestContext) SetTokenType(tokenType string) {
	d.TokenType = tokenType
}

func (d *DefaultIntrospectionRequestContext) IsActive() bool {
	return d.Active
}

func (d *DefaultIntrospectionRequestContext) SetActive(active bool) {
	d.Active = active
}

func (d *DefaultIntrospectionRequestContext) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultIntrospectionRequestContext) GetRequestedAt() time.Time {
	return d.RequestedAt
}

func (d *DefaultIntrospectionRequestContext) GetClientID() string {
	return d.ClientID
}

func (d *DefaultIntrospectionRequestContext) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultIntrospectionRequestContext) GetClientSecret() string {
	return d.Secret
}

func (d *DefaultIntrospectionRequestContext) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultIntrospectionRequestContext) GetError() sdk.IError {
	return d.Error
}

func (d *DefaultIntrospectionRequestContext) SetError(err sdk.IError) {
	d.Error = err
}

func (d *DefaultIntrospectionRequestContext) GetForm() *url.Values {
	return d.Form
}

func (d *DefaultIntrospectionRequestContext) GetProfile() sdk.RequestProfile {
	return d.Profile
}

func (d *DefaultIntrospectionRequestContext) SetProfile(profile sdk.RequestProfile) {
	d.Profile = profile
}

func (d *DefaultIntrospectionRequestContext) GetToken() string {
	return d.Token
}

func (d *DefaultIntrospectionRequestContext) GetTokenTypeHint() string {
	return d.TokenTypeHint
}
