package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"time"
)

type DefaultAuthCodeIssuer struct {
	AuthCodeStrategy sdk.IAuthorizationCodeStrategy
	Lifespan         time.Duration
}

func (d *DefaultAuthCodeIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetResponseType().Has(sdk.ResponseTypeCode) {
		expiry := requestContext.GetRequestedAt().Add(d.Lifespan).Round(time.Second)
		code, signature := d.AuthCodeStrategy.GenerateAuthCode()
		requestContext.IssueAuthorizationCode(code, signature, expiry)
	}
	return nil
}

func (d *DefaultAuthCodeIssuer) Configure(config *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IAuthorizationCodeStrategy); ok {
			d.AuthCodeStrategy = us
			break
		}
	}
	d.Lifespan = config.AuthCodeLifespan
	if d.AuthCodeStrategy == nil {
		panic("configuration failed for DefaultAuthCodeIssuer")
	}
}
