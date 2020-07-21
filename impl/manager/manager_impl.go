package manager

import (
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
)

type (
	DefaultManager struct {
		Config                              *sdk.Config
		TokenRequestContextFactory          sdk.TokenRequestContextFactory
		TokenResponseWriter                 sdk.TokenResponseWriter
		JsonErrorWriter                     sdk.JsonErrorWriter
		AuthenticationRequestContextFactory sdk.AuthenticationRequestContextFactory
		AuthenticationResponseWriter        sdk.AuthenticationResponseWriter
		RedirectErrorWriter                 sdk.RedirectErrorWriter
		AuthEPHandlers                      []sdk.IAuthEPHandler
		TokenEPHandlers                     []sdk.ITokenEPHandler
		IntrospectionEPHandlers             []sdk.IIntrospectionEPHandler
		RevocationEPHandlers                []sdk.IRevocationEPHandler
		ErrorStrategy                       sdk.ErrorStrategy
		UserSessionManager                  sdk.ISessionManager
		LoginPageHandler                    http.HandlerFunc
		ConsentPageHandler                  http.HandlerFunc
		IntrospectionRequestContextFactory  sdk.IntrospectionRequestContextFactory
		IntrospectionResponseWriter         sdk.IntrospectionResponseWriter
		RevocationRequestContextFactory     sdk.RevocationRequestContextFactory
		RevocationResponseWriter            sdk.RevocationResponseWriter
		SecretStore                         sdk.ISecretStore
	}
)
