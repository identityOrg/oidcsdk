package manager

import (
	"net/http"
	sdk "oauth2-oidc-sdk"
)

type (
	DefaultManager struct {
		Config                              *sdk.Config
		TokenRequestContextFactory          sdk.TokenRequestContextFactory
		TokenResponseWriter                 sdk.TokenResponseWriter
		TokenErrorWriter                    sdk.TokenErrorWriter
		AuthenticationRequestContextFactory sdk.AuthenticationRequestContextFactory
		AuthenticationResponseWriter        sdk.AuthenticationResponseWriter
		AuthenticationErrorWriter           sdk.AuthenticationErrorWriter
		AuthEPHandlers                      []sdk.IAuthEPHandler
		TokenEPHandlers                     []sdk.ITokenEPHandler
		ErrorStrategy                       sdk.ErrorStrategy
		UserSessionManager                  sdk.ISessionManager
		LoginPageHandler                    http.HandlerFunc
		ConsentPageHandler                  http.HandlerFunc
	}
)
