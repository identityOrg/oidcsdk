package manager

import (
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
)

type (
	DefaultManager struct {
		Config                  *sdk.Config
		RequestContextFactory   sdk.IRequestContextFactory
		ErrorWriter             sdk.IErrorWriter
		ResponseWriter          sdk.IResponseWriter
		AuthEPHandlers          []sdk.IAuthEPHandler
		TokenEPHandlers         []sdk.ITokenEPHandler
		IntrospectionEPHandlers []sdk.IIntrospectionEPHandler
		RevocationEPHandlers    []sdk.IRevocationEPHandler
		UserInfoEPHandlers      []sdk.IUserInfoEPHandler
		ErrorStrategy           sdk.ErrorStrategy
		UserSessionManager      sdk.ISessionManager
		LoginPageHandler        http.HandlerFunc
		ConsentPageHandler      http.HandlerFunc
		SecretStore             sdk.ISecretStore
	}
)
