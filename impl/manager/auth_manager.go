package manager

import (
	"net/http"
	sdk "oauth2-oidc-sdk"
)

func (d *DefaultManager) ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request) sdk.Result {
	panic("implement me")
}
