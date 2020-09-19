package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultTokenPersister struct {
	TokenStore            sdk.ITokenStore
	UserStore             sdk.IUserStore
	GlobalConsentRequired bool
}

func NewDefaultTokenPersister(tokenStore sdk.ITokenStore, userStore sdk.IUserStore, config *sdk.Config) *DefaultTokenPersister {
	return &DefaultTokenPersister{TokenStore: tokenStore, UserStore: userStore, GlobalConsentRequired: config.GlobalConsentRequired}
}

func (d *DefaultTokenPersister) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	tokens := requestContext.GetIssuedTokens()
	profile := requestContext.GetProfile()
	reqId := requestContext.GetRequestID()
	if d.GlobalConsentRequired {
		err := d.UserStore.StoreConsent(ctx, profile.GetUsername(), requestContext.GetClientID(), profile.GetScope())
		if err != nil {
			return sdkerror.ErrInvalidRequest //todo change error
		}
	}
	err := d.TokenStore.StoreTokenProfile(ctx, reqId, &tokens.TokenSignatures, profile)
	if err != nil {
		return sdkerror.ErrInvalidRequest //todo change error
	}
	return nil
}

func (d *DefaultTokenPersister) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	tokens := requestContext.GetIssuedTokens()
	profile := requestContext.GetProfile()
	reqId := requestContext.GetRequestID()
	err := d.TokenStore.StoreTokenProfile(ctx, reqId, &tokens.TokenSignatures, profile)
	if err != nil {
		return sdkerror.ErrInvalidRequest //todo change error
	}
	if requestContext.GetGrantType() == sdk.GrantRefreshToken {
		previousReqID := requestContext.GetPreviousRequestID()
		err := d.TokenStore.InvalidateWithRequestID(ctx, previousReqID, sdk.ExpireAccessToken|sdk.ExpireRefreshToken)
		if err != nil {
			return sdkerror.ErrInvalidGrant //todo correct it
		}
	}
	if requestContext.GetGrantType() == sdk.GrantAuthorizationCode {
		previousReqID := requestContext.GetPreviousRequestID()
		err := d.TokenStore.InvalidateWithRequestID(ctx, previousReqID, sdk.ExpireAuthorizationCode)
		if err != nil {
			return sdkerror.ErrInvalidGrant //todo correct it
		}
	}
	return nil
}
