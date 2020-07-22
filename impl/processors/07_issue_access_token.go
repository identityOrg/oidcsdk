package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"time"
)

type DefaultAccessTokenIssuer struct {
	AccessTokenStrategy sdk.IAccessTokenStrategy
	Lifespan            time.Duration
}

func (d *DefaultAccessTokenIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetResponseType().Has(sdk.ResponseTypeToken) {
		token, signature := d.AccessTokenStrategy.GenerateAccessToken()
		expiry := requestContext.GetRequestedAt().Add(d.Lifespan).Round(time.Second)
		requestContext.IssueAccessToken(token, signature, expiry)
	}
	return nil
}

func (d *DefaultAccessTokenIssuer) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	token, signature := d.AccessTokenStrategy.GenerateAccessToken()
	expiry := requestContext.GetRequestedAt().Add(d.Lifespan).Round(time.Second)
	requestContext.IssueAccessToken(token, signature, expiry)
	return nil
}

func (d *DefaultAccessTokenIssuer) Configure(config *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IAccessTokenStrategy); ok {
			d.AccessTokenStrategy = us
			break
		}
	}
	d.Lifespan = config.AccessTokenLifespan
}
