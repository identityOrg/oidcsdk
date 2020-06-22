package manager

import (
	sdk "oauth2-oidc-sdk"
)

type (
	DefaultManager struct {
		Config                        sdk.Config
		SessionFactory                sdk.SessionFactory
		TokensFactory                 sdk.TokensFactory
		TokenRequestFactory           sdk.TokenRequestFactory
		TokenResponseFactory          sdk.TokenResponseFactory
		TokenResponseWriter           sdk.TokenResponseWriter
		TokenErrorWriter              sdk.TokenErrorWriter
		AuthenticationRequestFactory  sdk.AuthenticationRequestFactory
		AuthenticationResponseFactory sdk.AuthenticationResponseFactory
		AuthenticationResponseWriter  sdk.AuthenticationResponseWriter
		AuthenticationErrorWriter     sdk.AuthenticationErrorWriter
		IDTokenStrategy               sdk.IIDTokenStrategy
		AccessTokenStrategy           sdk.IAccessTokenStrategy
		RefreshTokenStrategy          sdk.IRefreshTokenStrategy
		TokenStore                    sdk.ITokenStore
		UserStore                     sdk.IUserStore
		ClientStore                   sdk.IClientStore
		AuthEPHandlers                []sdk.IAuthEPHandler
		TokenEPHandlers               []sdk.ITokenEPHandler
		ErrorStrategy                 sdk.ErrorStrategy
	}
)
