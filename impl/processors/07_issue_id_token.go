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
	if openID && idTokenRequired {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		client := requestContext.GetClient()
		tokens := requestContext.GetIssuedTokens()
		var tClaims map[string]interface{}
		token, err := d.IDTokenStrategy.GenerateIDToken(ctx, profile, client, expiry, tClaims, tokens)
		if err != nil {
			return sdkerror.ErrMisconfiguration.WithHint(err.Error())
		}
		requestContext.IssueIDToken(token)
	}
	return nil
}

func (d *DefaultIDTokenIssuer) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	profile := requestContext.GetProfile()
	if profile.GetScope().Has(sdk.ScopeOpenid) && profile.GetGrantType() != sdk.GrantClientCredentials {
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		client := requestContext.GetClient()
		tokens := requestContext.GetIssuedTokens()
		var tClaims map[string]interface{}
		token, err := d.IDTokenStrategy.GenerateIDToken(ctx, profile, client, expiry, tClaims, tokens)
		if err != nil {
			return sdkerror.ErrMisconfiguration.WithHint(err.Error())
		}
		requestContext.IssueIDToken(token)
	}
	return nil
}
