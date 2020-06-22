package manager

import (
	"context"
	"net/http"
	sdk "oauth2-oidc-sdk"
)

func (d *DefaultManager) ProcessTokenEP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if tokenRequest, iError := d.TokenRequestFactory(r); iError != nil {
		d.handleTokenEPError(w, iError)
		return
	} else {
		if tokenResponse, iError := d.TokenResponseFactory(tokenRequest); iError != nil {
			d.handleTokenEPError(w, iError)
			return
		} else {
			for _, handler := range d.TokenEPHandlers {
				if iError := handler.HandleTokenEP(ctx, tokenRequest, tokenResponse); iError != nil {
					d.handleTokenEPError(w, iError)
				}
			}
			if err := d.TokenResponseWriter(tokenResponse, w); err != nil {
				d.ErrorStrategy(err, w)
			}
		}
	}
}

func (d *DefaultManager) handleTokenEPError(w http.ResponseWriter, iError sdk.IError) {
	err := d.TokenErrorWriter(iError, w)
	if err != nil {
		d.ErrorStrategy(err, w)
	}
}
