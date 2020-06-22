package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultAccessCodeValidator struct {
	TokenStore       sdk.ITokenStore
	AuthCodeStrategy sdk.IAuthorizationCodeStrategy
}

func (d *DefaultAccessCodeValidator) HandleTokenEP(_ context.Context, request sdk.ITokenRequest, response sdk.ITokenResponse) sdk.IError {
	if request.GetGrantType() == "authorization_code" {
		if request.GetAuthorizationCode() == "" {
			return sdkerror.InvalidGrant.WithDescription("'authorization_code' not provided")
		}

		authCode := request.GetAuthorizationCode()
		authCodeSignature := d.AuthCodeStrategy.SignAuthCode(authCode)
		if profile, err := d.TokenStore.GetProfileWithAuthCodeSign(authCodeSignature); err != nil {
			return sdkerror.InvalidGrant.WithDescription("invalid 'authorization_code'")
		} else {
			response.SetProfile(profile)
		}
	}
	return nil
}

func (d *DefaultAccessCodeValidator) Configure(strategy interface{}, _ sdk.Config, args ...interface{}) {
	d.AuthCodeStrategy = strategy.(sdk.IAuthorizationCodeStrategy)
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
			break
		}
	}
	if d.TokenStore == nil {
		panic("failed in init of DefaultAccessCodeValidator")
	}
}
