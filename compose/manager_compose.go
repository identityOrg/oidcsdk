package compose

import (
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/manager"
	"oauth2-oidc-sdk/impl/strategies"
	"oauth2-oidc-sdk/impl/tokenep"
	"oauth2-oidc-sdk/impl/tokens"
)

func DefaultManager(config sdk.Config, strategy interface{}, args ...interface{}) sdk.IManager {
	dManager := manager.DefaultManager{}
	dManager.Config = config
	dManager.TokenResponseFactory = tokenep.DefaultTokenResponseFactory
	dManager.TokenRequestFactory = tokenep.DefaultTokenRequestFactory
	dManager.TokenResponseWriter = tokenep.DefaultTokenResponseWriter
	dManager.TokenErrorWriter = tokenep.DefaultTokenErrorWriter

	dManager.TokensFactory = tokens.DefaultTokensFactory

	dManager.AuthenticationResponseFactory = nil
	dManager.AuthenticationRequestFactory = nil
	dManager.AuthenticationResponseWriter = nil
	dManager.AuthenticationErrorWriter = nil

	dManager.ErrorStrategy = strategies.DefaultLoggingErrorStrategy

	dManager.AccessTokenStrategy = strategy.(sdk.IAccessTokenStrategy)
	dManager.RefreshTokenStrategy = strategy.(sdk.IRefreshTokenStrategy)
	dManager.IDTokenStrategy = strategy.(sdk.IIDTokenStrategy)

	for _, arg := range args {
		if configurable, ok := arg.(sdk.Configurable); ok {
			configurable.Configure(&dManager, config, args)
		}
		if element, ok := arg.(sdk.IUserStore); ok {
			dManager.UserStore = element
		}
		if element, ok := arg.(sdk.IClientStore); ok {
			dManager.ClientStore = element
		}
		if element, ok := arg.(sdk.ITokenStore); ok {
			dManager.TokenStore = element
		}
		if element, ok := arg.(sdk.IIDTokenStrategy); ok {
			dManager.IDTokenStrategy = element
		}
		if element, ok := arg.(sdk.IAccessTokenStrategy); ok {
			dManager.AccessTokenStrategy = element
		}
		if element, ok := arg.(sdk.IRefreshTokenStrategy); ok {
			dManager.RefreshTokenStrategy = element
		}
		if element, ok := arg.(sdk.IAuthEPHandler); ok {
			dManager.AuthEPHandlers = append(dManager.AuthEPHandlers, element)
		}
		if element, ok := arg.(sdk.ITokenEPHandler); ok {
			dManager.TokenEPHandlers = append(dManager.TokenEPHandlers, element)
		}
	}

	return &dManager
}
