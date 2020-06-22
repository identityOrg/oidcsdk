package oauth2_oidc_sdk

import "context"

type (
	IAuthEPHandler interface {
		Handle(ctx context.Context, request IAuthenticationRequest, response IAuthenticationResponse) IError
	}

	ITokenEPHandler interface {
		Handle(ctx context.Context, request ITokenRequest, response ITokenResponse) IError
	}
)
