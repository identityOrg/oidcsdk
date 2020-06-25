package manager

import (
	"context"
	"net/http"
	sdk "oauth2-oidc-sdk"
)

func (d *DefaultManager) ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request) sdk.Result {
	ctx := context.Background()
	if authRequestContext, iError := d.AuthenticationRequestContextFactory(r); iError != nil {
		if authRequestContext != nil {
			authRequestContext.SetError(iError)
			err := d.AuthenticationErrorWriter(authRequestContext, w, r)
			if err != nil {
				d.ErrorStrategy(err, w)
			}
		} else {
			d.ErrorStrategy(iError, w)
		}
		return 0
	} else {
		for _, handler := range d.AuthEPHandlers {
			if iError := handler.HandleAuthEP(ctx, authRequestContext); iError != nil {
				authRequestContext.SetError(iError)
				err := d.AuthenticationErrorWriter(authRequestContext, w, r)
				if err != nil {
					d.ErrorStrategy(err, w)
				}
				return
			}
		}
		if err := d.AuthenticationResponseWriter(authRequestContext, w, r); err != nil {
			d.ErrorStrategy(err, w)
		}
	}
}
