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
	}
	ErrorStrategy func(err error, w http.ResponseWriter)
)
