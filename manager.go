package oauth2_oidc_sdk

import (
	"net/http"
)

type (
	IManager interface {
		ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request) Result
		ProcessTokenEP(w http.ResponseWriter, r *http.Request)
	}
	Result        uint8
	ErrorStrategy func(err error, w http.ResponseWriter)
)

const (
	ResultLoginRequired   Result = 1
	ResultConsentRequired Result = 2
)
