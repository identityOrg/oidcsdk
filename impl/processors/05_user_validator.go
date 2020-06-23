package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultUserValidator struct {
	UserStore sdk.IUserStore
}

func (d *DefaultUserValidator) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	grantType := requestContext.GetGrantType()
	if grantType == "password" {
		username := requestContext.GetUsername()
		password := requestContext.GetPassword()
		err := d.UserStore.Authenticate(username, []byte(password))
		if err != nil {
			return sdkerror.InvalidGrant.WithDescription("user authentication failed")
		}
		requestContext.GetProfile().SetUsername(username)
	} else if grantType == "client_credentials" {
		requestContext.GetProfile().SetUsername(requestContext.GetClientID())
	}
	return nil
}

func (d *DefaultUserValidator) Configure(_ interface{}, _ sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IUserStore); ok {
			d.UserStore = us
		}
	}
	if d.UserStore == nil {
		panic("failed to init DefaultUserValidator")
	}
}
