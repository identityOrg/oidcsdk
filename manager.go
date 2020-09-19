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

		SetLoginPageHandler(pageHandler http.HandlerFunc)
		SetConsentPageHandler(pageHandler http.HandlerFunc)
		SetErrorStrategy(strategy ErrorStrategy)
	}
	ErrorStrategy func(err error, w http.ResponseWriter)
)
