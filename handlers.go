package oidcsdk

import "context"

type (
	IAuthEPHandler interface {
		HandleAuthEP(ctx context.Context, requestContext IAuthenticationRequestContext) IError
	}

	ITokenEPHandler interface {
		HandleTokenEP(ctx context.Context, requestContext ITokenRequestContext) IError
	}

	IIntrospectionEPHandler interface {
		HandleIntrospectionEP(ctx context.Context, requestContext IIntrospectionRequestContext) IError
	}

	IRevocationEPHandler interface {
		HandleRevocationEP(ctx context.Context, requestContext IRevocationRequestContext) IError
	}

	IUserInfoEPHandler interface {
		HandleUserInfoEP(ctx context.Context, requestContext IUserInfoRequestContext) IError
	}
)
