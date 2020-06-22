package oauth2_oidc_sdk

import "context"

type (
	IAuthEPHandler interface {
		HandleAuthEP(ctx context.Context, request IAuthenticationRequest, response IAuthenticationResponse) IError
	}

	ITokenEPHandler interface {
		HandleTokenEP(ctx context.Context, request ITokenRequest, response ITokenResponse) IError
	}
)
