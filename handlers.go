package oauth2_oidc_sdk

import "context"

type (
	IAuthEPHandler interface {
		HandleAuthEP(ctx context.Context, requestContext IAuthenticationRequestContext) (IError, Result)
	}

	ITokenEPHandler interface {
		HandleTokenEP(ctx context.Context, requestContext ITokenRequestContext) IError
	}
)
