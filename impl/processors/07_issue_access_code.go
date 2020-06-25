package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"time"
)

type DefaultAuthCodeIssuer struct {
	AuthCodeStrategy sdk.IAuthorizationCodeStrategy
	Lifespan         time.Duration
}

func (d *DefaultAuthCodeIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) (sdk.IError, sdk.Result) {
	if requestContext.GetResponseType().Has("code") {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		code, signature := d.AuthCodeStrategy.GenerateAuthCode()
		requestContext.IssueAuthorizationCode(code, signature, expiry)
	}
	return nil, sdk.ResultNoOperation
}

func (d *DefaultAuthCodeIssuer) Configure(strategy interface{}, config *sdk.Config, _ ...interface{}) {
	d.AuthCodeStrategy = strategy.(sdk.IAuthorizationCodeStrategy)
	d.Lifespan = config.AuthCodeLifespan
}
