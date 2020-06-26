package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultUserValidator struct {
	UserStore   sdk.IUserStore
	ClientStore sdk.IClientStore
}

func (d *DefaultUserValidator) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) (sdk.IError, sdk.Result) {
	return sdkerror.ErrUnsupportedGrantType, sdk.ResultNoOperation
}

func (d *DefaultUserValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	grantType := requestContext.GetGrantType()
	if grantType == "password" {
		username := requestContext.GetUsername()
		password := requestContext.GetPassword()
		err := d.UserStore.Authenticate(username, []byte(password))
		if err != nil {
			return sdkerror.ErrInvalidGrant.WithDescription("user authentication failed")
		}
		profile := d.UserStore.FetchUserProfile(username)
		profile.SetScope(requestContext.GetGrantedScopes())
		profile.SetAudience(requestContext.GetGrantedAudience())
		requestContext.SetProfile(profile)
	} else if grantType == "client_credentials" {
		profile := d.ClientStore.FetchClientProfile(requestContext.GetClientID())
		profile.SetScope(requestContext.GetGrantedScopes())
		profile.SetAudience(requestContext.GetGrantedAudience())
		requestContext.SetProfile(profile)
	}
	return nil
}

func (d *DefaultUserValidator) Configure(_ interface{}, _ *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IUserStore); ok {
			d.UserStore = us
		}
		if cs, ok := arg.(sdk.IClientStore); ok {
			d.ClientStore = cs
		}
	}
	if d.UserStore == nil || d.ClientStore == nil {
		panic("failed to init DefaultUserValidator")
	}
}
