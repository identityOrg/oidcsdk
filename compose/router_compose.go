package compose

import (
	"github.com/gorilla/mux"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/middleware"
	"net/http"
)

func CreateNewRouter(sdkManager sdk.IManager) *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/oauth2").Subrouter()
	subRouter.Use(middleware.NoCache)
	subRouter.Methods(http.MethodPost).Path("/token").HandlerFunc(sdkManager.ProcessTokenEP)
	subRouter.Methods(http.MethodGet).Path("/authorize").HandlerFunc(sdkManager.ProcessAuthorizationEP)
	subRouter.Methods(http.MethodPost).Path("/introspect").HandlerFunc(sdkManager.ProcessIntrospectionEP)
	subRouter.Methods(http.MethodPost).Path("/revoke").HandlerFunc(sdkManager.ProcessRevocationEP)
	subRouter.Methods(http.MethodGet).Path("/keys").HandlerFunc(sdkManager.ProcessKeysEP)
	router.Methods(http.MethodGet).Path(sdk.UrlOidcDiscovery).Handler(middleware.NoCache(http.HandlerFunc(sdkManager.ProcessDiscoveryEP)))
	return router
}
