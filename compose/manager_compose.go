package compose

import (
	"net/http"
	sdk "oidcsdk"
	"oidcsdk/impl/authep"
	"oidcsdk/impl/manager"
	"oidcsdk/impl/strategies"
	"oidcsdk/impl/tokenep"
)

func DefaultManager(config *sdk.Config, strategy interface{}, args ...interface{}) sdk.IManager {
	dManager := manager.DefaultManager{}
	dManager.Config = config
	dManager.TokenRequestContextFactory = tokenep.DefaultTokenRequestContextFactory
	dManager.TokenResponseWriter = tokenep.DefaultTokenResponseWriter
	dManager.JsonErrorWriter = tokenep.DefaultJsonErrorWriter

	dManager.AuthenticationRequestContextFactory = authep.DefaultAuthenticationRequestContextFactory
	dManager.AuthenticationResponseWriter = authep.DefaultAuthenticationResponseWriter
	dManager.RedirectErrorWriter = authep.DefaultRedirectErrorWriter

	dManager.ErrorStrategy = strategies.DefaultLoggingErrorStrategy

	if configurable, ok := strategy.(sdk.IConfigurable); ok {
		configurable.Configure(strategy, config, args)
	}

	for _, arg := range args {
		if configurable, ok := arg.(sdk.IConfigurable); ok {
			configurable.Configure(strategy, config, args...)
		}
		if element, ok := arg.(sdk.IAuthEPHandler); ok {
			dManager.AuthEPHandlers = append(dManager.AuthEPHandlers, element)
		}
		if element, ok := arg.(sdk.ITokenEPHandler); ok {
			dManager.TokenEPHandlers = append(dManager.TokenEPHandlers, element)
		}
		if element, ok := arg.(sdk.ISessionManager); ok {
			dManager.UserSessionManager = element
		}
	}

	return &dManager
}

func SetLoginPageHandler(iManager sdk.IManager, handler http.HandlerFunc) {
	if defaultManager, ok := iManager.(*manager.DefaultManager); ok {
		defaultManager.LoginPageHandler = handler
	}
}

func SetConsentPageHandler(iManager sdk.IManager, handler http.HandlerFunc) {
	if defaultManager, ok := iManager.(*manager.DefaultManager); ok {
		defaultManager.ConsentPageHandler = handler
	}
}
