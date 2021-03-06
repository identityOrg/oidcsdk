package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultRefreshTokenValidator struct {
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
	TokenStore           sdk.ITokenStore
}

func NewDefaultRefreshTokenValidator(refreshTokenStrategy sdk.IRefreshTokenStrategy, tokenStore sdk.ITokenStore) *DefaultRefreshTokenValidator {
	return &DefaultRefreshTokenValidator{RefreshTokenStrategy: refreshTokenStrategy, TokenStore: tokenStore}
}

func (d *DefaultRefreshTokenValidator) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == sdk.GrantRefreshToken {
		if requestContext.GetRefreshToken() == "" {
			return sdkerror.ErrInvalidGrant.WithDescription("'refresh_token' not provided")
		}

		refreshToken := requestContext.GetRefreshToken()
		refreshTokenSignature, err := d.RefreshTokenStrategy.SignRefreshToken(refreshToken)
		if err != nil {
			return sdkerror.ErrInvalidGrant.WithDescription("invalid 'refresh_token'")
		}
		if profile, reqId, err := d.TokenStore.GetProfileWithRefreshTokenSign(ctx, refreshTokenSignature); err != nil {
			return sdkerror.ErrInvalidGrant.WithDescription("invalid 'refresh_token'")
		} else {
			requestContext.SetProfile(profile)
			requestContext.SetPreviousRequestID(reqId)
		}
	}
	return nil
}
