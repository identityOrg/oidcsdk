package manager

import (
	"net/http"
)

func (d *DefaultManager) ProcessTokenEP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if tokenRequestContext, iError := d.RequestContextFactory.BuildTokenRequestContext(r); iError != nil {
		err := d.ErrorWriter.WriteJsonError(iError, nil, w, r)
		if err != nil {
			d.ErrorStrategy(err, w)
		}
		return
	} else {
		for _, handler := range d.TokenEPHandlers {
			if iError := handler.HandleTokenEP(ctx, tokenRequestContext); iError != nil {
				err := d.ErrorWriter.WriteJsonError(iError, nil, w, r)
				if err != nil {
					d.ErrorStrategy(err, w)
				}
				return
			}
		}
		if err := d.ResponseWriter.WriteTokenResponse(tokenRequestContext, w, r); err != nil {
			d.ErrorStrategy(err, w)
		}
	}
}
