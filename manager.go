package oauth2_oidc_sdk

import (
	"net/http"
)

type (
	IManager interface {
		ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request)
		ProcessTokenEP(w http.ResponseWriter, r *http.Request)
	}
	Result        uint8
	ErrorStrategy func(err error, w http.ResponseWriter)
)

const (
	ResultNoOperation     Result = 1
	ResultLoginRequired   Result = 2
	ResultConsentRequired Result = 4
)
