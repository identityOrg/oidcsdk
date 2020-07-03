package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
	"time"
)

type DefaultIDTokenIssuer struct {
	IDTokenStrategy sdk.IIDTokenStrategy
	Lifespan        time.Duration
}

func (d *DefaultIDTokenIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetResponseType().Has(sdk.ResponseTypeIdToken) {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		profile := requestContext.GetProfile()
		client := requestContext.GetClient()
		var tClaims map[string]interface{}
		token, err := d.IDTokenStrategy.GenerateIDToken(profile, client, expiry, tClaims)
		if err != nil {
			return sdkerror.ErrMisconfiguration.WithHint(err.Error())
		}
		requestContext.IssueIDToken(token)
	}
	return nil
}

func (d *DefaultIDTokenIssuer) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid) {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		profile := requestContext.GetProfile()
		client := requestContext.GetClient()
		var tClaims map[string]interface{}
		token, err := d.IDTokenStrategy.GenerateIDToken(profile, client, expiry, tClaims)
		if err != nil {
			return sdkerror.ErrMisconfiguration.WithHint(err.Error())
		}
		requestContext.IssueIDToken(token)
	}
	return nil
}

func (d *DefaultIDTokenIssuer) Configure(strategy interface{}, config *sdk.Config, _ ...interface{}) {
	d.IDTokenStrategy = strategy.(sdk.IIDTokenStrategy)
	d.Lifespan = config.AccessTokenLifespan
}
