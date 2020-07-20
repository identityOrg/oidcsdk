package manager

import (
	"net/http"
	sdk "oidcsdk"
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
		ErrorStrategy                       sdk.ErrorStrategy
		UserSessionManager                  sdk.ISessionManager
		LoginPageHandler                    http.HandlerFunc
		ConsentPageHandler                  http.HandlerFunc
	}
)
