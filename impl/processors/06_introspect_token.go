package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultTokenIntrospectionProcessor struct {
	TokenStore           sdk.ITokenStore
	AccessTokenStrategy  sdk.IAccessTokenStrategy
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
}

func NewDefaultTokenIntrospectionProcessor(tokenStore sdk.ITokenStore, accessTokenStrategy sdk.IAccessTokenStrategy, refreshTokenStrategy sdk.IRefreshTokenStrategy) *DefaultTokenIntrospectionProcessor {
	return &DefaultTokenIntrospectionProcessor{TokenStore: tokenStore, AccessTokenStrategy: accessTokenStrategy, RefreshTokenStrategy: refreshTokenStrategy}
}

func (d *DefaultTokenIntrospectionProcessor) HandleIntrospectionEP(ctx context.Context, requestContext sdk.IIntrospectionRequestContext) sdk.IError {
	if requestContext.GetToken() == "" {
		return sdkerror.ErrInvalidRequest.WithHint("token is not provided or blank")
	}
	requestContext.SetActive(false)
	switch requestContext.GetTokenTypeHint() {
	case "":
		fallthrough
	case "access_token":
		fallthrough
	case "bearer":
		tokenSign, err := d.AccessTokenStrategy.SignAccessToken(requestContext.GetToken())
		if err != nil {
			return sdkerror.ErrInvalidRequest.WithHint(err.Error())
		}
		profile, _, err := d.TokenStore.GetProfileWithAccessTokenSign(ctx, tokenSign)
		if err == nil {
			requestContext.SetActive(true)
			requestContext.SetProfile(profile)
			return nil
		}
		fallthrough
	case "refresh_token":
		tokenSign, err := d.RefreshTokenStrategy.SignRefreshToken(requestContext.GetToken())
		if err != nil {
			return sdkerror.ErrInvalidRequest.WithHint(err.Error())
		}
		profile, _, err := d.TokenStore.GetProfileWithRefreshTokenSign(ctx, tokenSign)
		if err == nil {
			requestContext.SetActive(true)
			requestContext.SetProfile(profile)
			return nil
		}
	case "id_token":
		fallthrough
	default:
		return sdkerror.ErrUnknownRequest.WithHint("token type hint unknown")
	}

	return nil
}
