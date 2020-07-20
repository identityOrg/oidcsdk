package manager

import (
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"net/http"
)

func (d *DefaultManager) ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if authRequestContext, iError := d.AuthenticationRequestContextFactory(r); iError != nil {
		if authRequestContext != nil {
			authRequestContext.SetError(iError)
			err := d.RedirectErrorWriter(authRequestContext, w, r)
			if err != nil {
				d.ErrorStrategy(err, w)
			}
		} else {
			d.ErrorStrategy(iError, w)
		}
	} else {
		if sess, err := d.UserSessionManager.RetrieveUserSession(r); err == nil {
			authRequestContext.SetUserSession(sess)
		}
		for _, handler := range d.AuthEPHandlers {
			if iError := handler.HandleAuthEP(ctx, authRequestContext); iError != nil {
				if iError.Error() == sdkerror.ErrLoginRequired.Name && iError.GetReason() == "" {
					d.LoginPageHandler(w, r)
				}
				if iError.Error() == sdkerror.ErrConsentRequired.Name && iError.GetReason() == "" {
					d.ConsentPageHandler(w, r)
				}
				authRequestContext.SetError(iError)
				err := d.RedirectErrorWriter(authRequestContext, w, r)
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
