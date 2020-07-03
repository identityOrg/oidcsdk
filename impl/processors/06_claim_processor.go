package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultClaimProcessor struct {
	UserStore sdk.IUserStore
}

func (d *DefaultClaimProcessor) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	openidReq := requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid)
	if openidReq && !requestContext.GetProfile().IsClient() {
		claims, err := d.UserStore.GetClaims(ctx, requestContext.GetProfile().GetUsername(), requestContext.GetProfile().GetScope(), nil)
		if err != nil {
			return sdkerror.ErrServerError.WithHint(err.Error())
		}
		contextClaim := requestContext.GetClaims()
		for s, i := range claims {
			contextClaim[s] = i
		}
	}
	return nil
}

func (d *DefaultClaimProcessor) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid) {
		if requestContext.GetResponseType().Has(sdk.ResponseTypeIdToken) {
			claims, err := d.UserStore.GetClaims(ctx, requestContext.GetProfile().GetUsername(), requestContext.GetProfile().GetScope(), nil)
			if err != nil {
				return sdkerror.ErrServerError.WithHint(err.Error())
			}
			contextClaim := requestContext.GetClaims()
			for s, i := range claims {
				contextClaim[s] = i
			}
		}
	}
	return nil
}

func (d *DefaultClaimProcessor) Configure(_ interface{}, _ *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IUserStore); ok {
			d.UserStore = us
			break
		}
	}
	if d.UserStore == nil {
		panic("failed to initialize DefaultClaimProcessor")
	}
}
