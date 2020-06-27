package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultRefreshTokenValidator struct {
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
	TokenStore           sdk.ITokenStore
}

func (d *DefaultRefreshTokenValidator) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == "refresh_token" {
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

func (d *DefaultRefreshTokenValidator) Configure(strategy interface{}, _ *sdk.Config, args ...interface{}) {
	d.RefreshTokenStrategy = strategy.(sdk.IRefreshTokenStrategy)
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
			break
		}
	}
	if d.TokenStore == nil {
		panic("failed in init of DefaultRefreshTokenValidator")
	}
}
