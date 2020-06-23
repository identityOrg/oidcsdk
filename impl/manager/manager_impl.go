package manager

import (
	sdk "oauth2-oidc-sdk"
)

type (
	DefaultManager struct {
		Config                              sdk.Config
		TokenRequestContextFactory          sdk.TokenRequestContextFactory
		TokenResponseWriter                 sdk.TokenResponseWriter
		TokenErrorWriter                    sdk.TokenErrorWriter
		AuthenticationRequestContextFactory sdk.AuthenticationRequestContextFactory
		AuthenticationResponseWriter        sdk.AuthenticationResponseWriter
		AuthenticationErrorWriter           sdk.AuthenticationErrorWriter
		IDTokenStrategy                     sdk.IIDTokenStrategy
		AccessTokenStrategy                 sdk.IAccessTokenStrategy
		RefreshTokenStrategy                sdk.IRefreshTokenStrategy
		TokenStore                          sdk.ITokenStore
		UserStore                           sdk.IUserStore
		ClientStore                         sdk.IClientStore
		AuthEPHandlers                      []sdk.IAuthEPHandler
		TokenEPHandlers                     []sdk.ITokenEPHandler
		ErrorStrategy                       sdk.ErrorStrategy
	}
)
