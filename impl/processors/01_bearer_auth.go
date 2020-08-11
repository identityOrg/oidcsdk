package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultBearerUserAuthProcessor struct {
	TokenStore          sdk.ITokenStore
	UserStore           sdk.IUserStore
	AccessTokenStrategy sdk.IAccessTokenStrategy
}

func (d *DefaultBearerUserAuthProcessor) Configure(config *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
		}
		if us, ok := arg.(sdk.IUserStore); ok {
			d.UserStore = us
		}
		if ats, ok := arg.(sdk.IAccessTokenStrategy); ok {
			d.AccessTokenStrategy = ats
		}
	}
	if d.TokenStore == nil || d.UserStore == nil || d.AccessTokenStrategy == nil {
		panic("DefaultBearerUserAuthProcessor configuration failed")
	}
}

func (d *DefaultBearerUserAuthProcessor) HandleUserInfoEP(ctx context.Context, requestContext sdk.IUserInfoRequestContext) sdk.IError {
	token := requestContext.GetBearerToken()
	signature, err := d.AccessTokenStrategy.SignAccessToken(token)
	if err != nil {
		return sdkerror.ErrAccessDenied.WithHintf("invalid access token")
	}
	profile, _, err := d.TokenStore.GetProfileWithAccessTokenSign(ctx, signature)
	if err != nil {
		return sdkerror.ErrAccessDenied.WithHintf("invalid access token")
	}
	requestContext.SetUsername(profile.GetUsername())
	requestContext.SetApprovedScopes(profile.GetScope())
	//requestContext.SetRequestedClaims() // if claims parameter implemented
	return nil
}
