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

func NewDefaultAuthCodeIssuer(authCodeStrategy sdk.IAuthorizationCodeStrategy, config *sdk.Config) *DefaultAuthCodeIssuer {
	return &DefaultAuthCodeIssuer{AuthCodeStrategy: authCodeStrategy, Lifespan: config.AuthCodeLifespan}
}

func (d *DefaultAuthCodeIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetResponseType().Has(sdk.ResponseTypeCode) {
		expiry := requestContext.GetRequestedAt().Add(d.Lifespan).Round(time.Second)
		code, signature := d.AuthCodeStrategy.GenerateAuthCode()
		requestContext.IssueAuthorizationCode(code, signature, expiry)
	}
	return nil
}
