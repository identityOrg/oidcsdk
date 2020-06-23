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

func (d *DefaultRefreshTokenValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == "refresh_token" {
		if requestContext.GetRefreshToken() == "" {
			return sdkerror.InvalidGrant.WithDescription("'refresh_token' not provided")
		}

		refreshToken := requestContext.GetAuthorizationCode()
		refreshTokenSignature := d.RefreshTokenStrategy.SignRefreshToken(refreshToken)
		if profile, err := d.TokenStore.GetProfileWithRefreshTokenSign(refreshTokenSignature); err != nil {
			return sdkerror.InvalidGrant.WithDescription("invalid 'refresh_token'")
		} else {
			requestContext.SetProfile(profile)
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
