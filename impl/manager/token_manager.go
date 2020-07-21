package manager

import (
	"net/http"
)

func (d *DefaultManager) ProcessTokenEP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if tokenRequestContext, iError := d.TokenRequestContextFactory(r); iError != nil {
		err := d.JsonErrorWriter(iError, w, r)
		if err != nil {
			d.ErrorStrategy(err, w)
		}
		return
	} else {
		for _, handler := range d.TokenEPHandlers {
			if iError := handler.HandleTokenEP(ctx, tokenRequestContext); iError != nil {
				err := d.JsonErrorWriter(iError, w, r)
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
