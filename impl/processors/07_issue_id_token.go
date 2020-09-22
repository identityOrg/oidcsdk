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

func NewDefaultIDTokenIssuer(IDTokenStrategy sdk.IIDTokenStrategy, config *sdk.Config) *DefaultIDTokenIssuer {
	return &DefaultIDTokenIssuer{IDTokenStrategy: IDTokenStrategy, Lifespan: config.AccessTokenLifespan}
}

func (d *DefaultIDTokenIssuer) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	idTokenRequired := requestContext.GetResponseType().Has(sdk.ResponseTypeIdToken)
	openID := requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid)
	profile := requestContext.GetProfile()
	profile.SetNonce(requestContext.GetNonce())
	if openID && idTokenRequired {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		client := requestContext.GetClient()
		var tClaims map[string]interface{}
		token, err := d.IDTokenStrategy.GenerateIDToken(ctx, profile, client, expiry, tClaims)
		if err != nil {
			return sdkerror.ErrMisconfiguration.WithHint(err.Error())
		}
		requestContext.IssueIDToken(token)
	}
	return nil
}

func (d *DefaultIDTokenIssuer) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid) {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		profile := requestContext.GetProfile()
		client := requestContext.GetClient()
		var tClaims map[string]interface{}
		token, err := d.IDTokenStrategy.GenerateIDToken(ctx, profile, client, expiry, tClaims)
		if err != nil {
			return sdkerror.ErrMisconfiguration.WithHint(err.Error())
		}
		requestContext.IssueIDToken(token)
	}
	return nil
}
