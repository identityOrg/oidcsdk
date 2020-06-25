package manager

import (
	"context"
	"net/http"
)

func (d *DefaultManager) ProcessTokenEP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if tokenRequestContext, iError := d.TokenRequestContextFactory(r); iError != nil {
		if tokenRequestContext != nil {
			tokenRequestContext.SetError(iError)
			err := d.TokenErrorWriter(tokenRequestContext, w, r)
			if err != nil {
				d.ErrorStrategy(err, w)
			}
		} else {
			d.ErrorStrategy(iError, w)
		}
		return
	} else {
		for _, handler := range d.TokenEPHandlers {
			if iError := handler.HandleTokenEP(ctx, tokenRequestContext); iError != nil {
				tokenRequestContext.SetError(iError)
				err := d.TokenErrorWriter(tokenRequestContext, w, r)
				if err != nil {
					d.ErrorStrategy(err, w)
				}
				return
			}
		}
		if err := d.TokenResponseWriter(tokenRequestContext, w, r); err != nil {
			d.ErrorStrategy(err, w)
		}
	}
}
