package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
)

type DefaultAccessCodeIssuer struct {
	AuthorizationCodeStrategy sdk.IAuthorizationCodeStrategy
	TokenStore                sdk.ITokenStore
}

func (d *DefaultAccessCodeIssuer) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	//profile := response.GetProfile()
	//code, signature := d.AuthorizationCodeStrategy.GenerateAuthCode()
	//d.TokenStore.
	return nil
}

func (d *DefaultAccessCodeIssuer) Configure(strategy interface{}, config sdk.Config, args ...interface{}) {
	d.AuthorizationCodeStrategy = strategy.(sdk.IAuthorizationCodeStrategy)
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
		}
	}
	if d.TokenStore == nil {
		panic("failed to initialize DefaultAccessCodeIssuer")
	}
}
