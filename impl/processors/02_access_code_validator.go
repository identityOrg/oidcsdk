package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultAccessCodeValidator struct {
	TokenStore       sdk.ITokenStore
	AuthCodeStrategy sdk.IAuthorizationCodeStrategy
}

func NewDefaultAccessCodeValidator(tokenStore sdk.ITokenStore, authCodeStrategy sdk.IAuthorizationCodeStrategy) *DefaultAccessCodeValidator {
	return &DefaultAccessCodeValidator{TokenStore: tokenStore, AuthCodeStrategy: authCodeStrategy}
}

func (d *DefaultAccessCodeValidator) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == sdk.GrantAuthorizationCode {
		if requestContext.GetAuthorizationCode() == "" {
			return sdkerror.ErrInvalidRequest.WithDescription("'authorization_code' not provided")
		}

		authCode := requestContext.GetAuthorizationCode()
		authCodeSignature, err := d.AuthCodeStrategy.SignAuthCode(authCode)
		if err != nil {
			return sdkerror.ErrInvalidGrant.WithDescription("invalid 'authorization_code'")
		}
		if profile, reqId, err := d.TokenStore.GetProfileWithAuthCodeSign(ctx, authCodeSignature); err != nil {
			return sdkerror.ErrInvalidGrant.WithDescription("invalid 'authorization_code'")
		} else {
			requestContext.SetProfile(profile)
			requestContext.SetPreviousRequestID(reqId)
		}
	}
	return nil
}
