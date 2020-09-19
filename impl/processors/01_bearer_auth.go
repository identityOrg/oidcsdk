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

func NewDefaultBearerUserAuthProcessor(tokenStore sdk.ITokenStore, userStore sdk.IUserStore, accessTokenStrategy sdk.IAccessTokenStrategy) *DefaultBearerUserAuthProcessor {
	return &DefaultBearerUserAuthProcessor{TokenStore: tokenStore, UserStore: userStore, AccessTokenStrategy: accessTokenStrategy}
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
