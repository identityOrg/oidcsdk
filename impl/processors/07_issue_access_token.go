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

func NewDefaultAccessTokenIssuer(accessTokenStrategy sdk.IAccessTokenStrategy, config *sdk.Config) *DefaultAccessTokenIssuer {
	return &DefaultAccessTokenIssuer{AccessTokenStrategy: accessTokenStrategy, Lifespan: config.AccessTokenLifespan}
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
