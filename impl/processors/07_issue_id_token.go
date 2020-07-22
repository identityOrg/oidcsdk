package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"time"
)

type DefaultIDTokenIssuer struct {
	IDTokenStrategy sdk.IIDTokenStrategy
	Lifespan        time.Duration
}

func (d *DefaultIDTokenIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	idTokenRequired := requestContext.GetResponseType().Has(sdk.ResponseTypeIdToken)
	openID := requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid)
	if openID && idTokenRequired {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		profile := requestContext.GetProfile()
		client := requestContext.GetClient()
		nonce := requestContext.GetNonce()
		profile.SetNonce(nonce)
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

func (d *DefaultIDTokenIssuer) Configure(config *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IIDTokenStrategy); ok {
			d.IDTokenStrategy = us
			break
		}
	}
	if d.IDTokenStrategy == nil {
		panic("failed to initialize DefaultIDTokenIssuer")
	}
	d.Lifespan = config.AccessTokenLifespan
}
