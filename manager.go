package oidcsdk

import (
	"net/http"
)

type (
	IManager interface {
		ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request)
		ProcessTokenEP(w http.ResponseWriter, r *http.Request)
		ProcessIntrospectionEP(w http.ResponseWriter, r *http.Request)
		ProcessRevocationEP(w http.ResponseWriter, r *http.Request)
	}
	ErrorStrategy func(err error, w http.ResponseWriter)
)
