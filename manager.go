package oidcsdk

import (
	"net/http"
)

type (
	IManager interface {
		ProcessAuthorizationEP(writer http.ResponseWriter, request *http.Request)
		ProcessTokenEP(writer http.ResponseWriter, request *http.Request)
		ProcessIntrospectionEP(writer http.ResponseWriter, request *http.Request)
		ProcessRevocationEP(writer http.ResponseWriter, request *http.Request)
		ProcessDiscoveryEP(writer http.ResponseWriter, request *http.Request)
		ProcessKeysEP(writer http.ResponseWriter, request *http.Request)
		ProcessUserInfoEP(writer http.ResponseWriter, request *http.Request)
		ProcessRPILogoutEP(writer http.ResponseWriter, request *http.Request)

		SetErrorStrategy(strategy ErrorStrategy)
	}
	ErrorStrategy        func(err error, w http.ResponseWriter)
	IPageResponseHandler interface {
		DisplayLogoutConsentPage(w http.ResponseWriter, r *http.Request)
		DisplayLogoutStatusPage(w http.ResponseWriter, r *http.Request)
		DisplayErrorPage(err error, w http.ResponseWriter, r *http.Request)
		DisplayLoginPage(w http.ResponseWriter, r *http.Request)
		DisplayConsentPage(w http.ResponseWriter, r *http.Request)
	}
)
