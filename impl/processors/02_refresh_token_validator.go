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

func (d *DefaultRefreshTokenValidator) HandleTokenEP(_ context.Context, request sdk.ITokenRequest, response sdk.ITokenResponse) sdk.IError {
	if request.GetGrantType() == "refresh_token" {
		if request.GetRefreshToken() == "" {
			return sdkerror.InvalidGrant.WithDescription("'refresh_token' not provided")
		}

		refreshToken := request.GetAuthorizationCode()
		refreshTokenSignature := d.RefreshTokenStrategy.SignRefreshToken(refreshToken)
		if profile, err := d.TokenStore.GetProfileWithRefreshTokenSign(refreshTokenSignature); err != nil {
			return sdkerror.InvalidGrant.WithDescription("invalid 'refresh_token'")
		} else {
			response.SetProfile(profile)
		}
	}
	return nil
}

func (d *DefaultRefreshTokenValidator) Configure(strategy interface{}, config sdk.Config, args ...interface{}) {
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
