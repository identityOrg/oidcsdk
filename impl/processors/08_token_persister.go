package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultTokenPersister struct {
	TokenStore sdk.ITokenStore
}

func (d *DefaultTokenPersister) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	tokens := requestContext.GetIssuedTokens()
	profile := requestContext.GetProfile()
	reqId := requestContext.GetRequestID()
	err := d.TokenStore.StoreTokenProfile(reqId, tokens.TokenSignatures, profile)
	if err != nil {
		return sdkerror.InvalidRequest //todo change error
	}
	return nil
}

func (d *DefaultTokenPersister) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	tokens := requestContext.GetIssuedTokens()
	profile := requestContext.GetProfile()
	reqId := requestContext.GetRequestID()
	err := d.TokenStore.StoreTokenProfile(reqId, tokens.TokenSignatures, profile)
	if err != nil {
		return sdkerror.InvalidRequest //todo change error
	}
	return nil
}

func (d *DefaultTokenPersister) Configure(_ interface{}, _ *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
		}
	}
}
