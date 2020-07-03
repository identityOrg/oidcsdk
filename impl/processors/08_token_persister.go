package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultTokenPersister struct {
	TokenStore           sdk.ITokenStore
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
}

func (d *DefaultTokenPersister) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	tokens := requestContext.GetIssuedTokens()
	profile := requestContext.GetProfile()
	reqId := requestContext.GetRequestID()
	err := d.TokenStore.StoreTokenProfile(ctx, reqId, tokens.TokenSignatures, profile)
	if err != nil {
		return sdkerror.ErrInvalidRequest //todo change error
	}
	return nil
}

func (d *DefaultTokenPersister) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	tokens := requestContext.GetIssuedTokens()
	profile := requestContext.GetProfile()
	reqId := requestContext.GetRequestID()
	err := d.TokenStore.StoreTokenProfile(ctx, reqId, tokens.TokenSignatures, profile)
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

func (d *DefaultTokenPersister) Configure(strategy interface{}, _ *sdk.Config, args ...interface{}) {
	d.RefreshTokenStrategy = strategy.(sdk.IRefreshTokenStrategy)
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
		}
	}
	if d.TokenStore == nil {
		panic("failed to init DefaultTokenPersister")
	}
}
