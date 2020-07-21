package compose

import (
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/authep"
	"github.com/identityOrg/oidcsdk/impl/introrevoke"
	"github.com/identityOrg/oidcsdk/impl/manager"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"github.com/identityOrg/oidcsdk/impl/tokenep"
	"net/http"
)

func DefaultManager(config *sdk.Config, args ...interface{}) sdk.IManager {
	dManager := manager.DefaultManager{}
	dManager.Config = config
	dManager.TokenRequestContextFactory = tokenep.DefaultTokenRequestContextFactory
	dManager.TokenResponseWriter = tokenep.DefaultTokenResponseWriter
	dManager.JsonErrorWriter = tokenep.DefaultJsonErrorWriter

	dManager.AuthenticationRequestContextFactory = authep.DefaultAuthenticationRequestContextFactory
	dManager.AuthenticationResponseWriter = authep.DefaultAuthenticationResponseWriter
	dManager.RedirectErrorWriter = authep.DefaultRedirectErrorWriter

	dManager.IntrospectionRequestContextFactory = introrevoke.DefaultIntrospectionRequestContextFactory
	dManager.IntrospectionResponseWriter = introrevoke.DefaultIntrospectionResponseWriter

	dManager.RevocationRequestContextFactory = introrevoke.DefaultRevocationRequestContextFactory
	dManager.RevocationResponseWriter = introrevoke.DefaultRevocationResponseWriter

	dManager.ErrorStrategy = strategies.DefaultLoggingErrorStrategy

	for _, arg := range args {
		if configurable, ok := arg.(sdk.IConfigurable); ok {
			configurable.Configure(config, args...)
		}
		if element, ok := arg.(sdk.IAuthEPHandler); ok {
			dManager.AuthEPHandlers = append(dManager.AuthEPHandlers, element)
		}
		if element, ok := arg.(sdk.ITokenEPHandler); ok {
			dManager.TokenEPHandlers = append(dManager.TokenEPHandlers, element)
		}
		if element, ok := arg.(sdk.IIntrospectionEPHandler); ok {
			dManager.IntrospectionEPHandlers = append(dManager.IntrospectionEPHandlers, element)
		}
		if element, ok := arg.(sdk.IRevocationEPHandler); ok {
			dManager.RevocationEPHandlers = append(dManager.RevocationEPHandlers, element)
		}
		if element, ok := arg.(sdk.ISessionManager); ok {
			dManager.UserSessionManager = element
		}
		if element, ok := arg.(sdk.ISecretStore); ok {
			dManager.SecretStore = element
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
