package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultClaimProcessor struct {
	UserStore sdk.IUserStore
}

func NewDefaultClaimProcessor(userStore sdk.IUserStore) *DefaultClaimProcessor {
	return &DefaultClaimProcessor{UserStore: userStore}
}

func (d *DefaultClaimProcessor) HandleUserInfoEP(ctx context.Context, requestContext sdk.IUserInfoRequestContext) sdk.IError {
	username := requestContext.GetUsername()
	scopes := requestContext.GetApprovedScopes()
	claimIds := requestContext.GetRequestedClaims()
	claims, err := d.UserStore.GetClaims(ctx, username, scopes, claimIds)
	if err != nil {
		return sdkerror.ErrServerError.WithHint("failed to fetch user claims").WithDebug(err.Error())
	}
	for s, _ := range claims {
		requestContext.AddClaim(s, claims[s])
	}
	return nil
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
