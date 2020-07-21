package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultTokenRevocationProcessor struct {
	TokenStore           sdk.ITokenStore
	AccessTokenStrategy  sdk.IAccessTokenStrategy
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
}

func (d *DefaultTokenRevocationProcessor) HandleRevocationEP(ctx context.Context, requestContext sdk.IRevocationRequestContext) sdk.IError {
	if requestContext.GetToken() == "" {
		return nil // no error needed
	}

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
		profile, reqID, err := d.TokenStore.GetProfileWithAccessTokenSign(ctx, tokenSign)
		if err == nil && reqID != "" {
			if profile.GetClientID() != requestContext.GetClientID() {
				return sdkerror.ErrRevokationClientMismatch.WithHintf("client id mismatch token vs request")
			}
			err = d.TokenStore.InvalidateWithRequestID(ctx, reqID, sdk.ExpireAccessToken)
			if err != nil {
				return sdkerror.ErrServerError.WithHint(err.Error())
			}
			return nil
		}
		fallthrough
	case "refresh_token":
		tokenSign, err := d.RefreshTokenStrategy.SignRefreshToken(requestContext.GetToken())
		if err != nil {
			return sdkerror.ErrInvalidRequest.WithHint(err.Error())
		}
		profile, reqID, err := d.TokenStore.GetProfileWithRefreshTokenSign(ctx, tokenSign)
		if err == nil && reqID != "" {
			if profile.GetClientID() != requestContext.GetClientID() {
				return sdkerror.ErrRevokationClientMismatch.WithHintf("client id mismatch token vs request")
			}
			err = d.TokenStore.InvalidateWithRequestID(ctx, reqID, sdk.ExpireRefreshToken|sdk.ExpireAccessToken)
			if err != nil {
				return sdkerror.ErrServerError.WithHint(err.Error())
			}
			return nil
		}
	default:
		return sdkerror.ErrUnsupportedTokenType.WithHint("token type not supported for revocation")
	}

	return nil
}

func (d *DefaultTokenRevocationProcessor) Configure(_ *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if store, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = store
		}
		if ts, ok := arg.(sdk.IAccessTokenStrategy); ok {
			d.AccessTokenStrategy = ts
		}
		if us, ok := arg.(sdk.IRefreshTokenStrategy); ok {
			d.RefreshTokenStrategy = us
		}
	}
	if d.TokenStore == nil || d.AccessTokenStrategy == nil || d.RefreshTokenStrategy == nil {
		panic("DefaultTokenRevocationProcessor configuration failed")
	}
}
