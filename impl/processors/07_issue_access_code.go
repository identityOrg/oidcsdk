package processors

import (
	"context"
	sdk "oidcsdk"
	"time"
)

type DefaultAuthCodeIssuer struct {
	AuthCodeStrategy sdk.IAuthorizationCodeStrategy
	Lifespan         time.Duration
}

func (d *DefaultAuthCodeIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetResponseType().Has(sdk.ResponseTypeCode) {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		code, signature := d.AuthCodeStrategy.GenerateAuthCode()
		requestContext.IssueAuthorizationCode(code, signature, expiry)
	}
	return nil
}

func (d *DefaultAuthCodeIssuer) Configure(strategy interface{}, config *sdk.Config, _ ...interface{}) {
	d.AuthCodeStrategy = strategy.(sdk.IAuthorizationCodeStrategy)
	d.Lifespan = config.AuthCodeLifespan
}
