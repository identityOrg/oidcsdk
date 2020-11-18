package manager

import (
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"net/http"
)

func (d *DefaultManager) ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if authRequestContext, iError := d.RequestContextFactory.BuildAuthorizationRequestContext(r); iError != nil {
		if authRequestContext != nil {
			authRequestContext.SetError(iError)
			err := d.ErrorWriter.WriteRedirectError(authRequestContext, w, r)
			if err != nil {
				d.ErrorStrategy(err, w)
			}
		} else {
			d.ErrorStrategy(iError, w)
		}
	} else {
		if sess, err := d.UserSessionManager.RetrieveUserSession(w, r); err == nil {
			authRequestContext.SetUserSession(sess)
		}
		for _, handler := range d.AuthEPHandlers {
			if iError := handler.HandleAuthEP(ctx, authRequestContext); iError != nil {
				if iError.Error() == sdkerror.ErrLoginRequired.Name && iError.GetReason() == "" {
					d.PageResponseHandler.DisplayLoginPage(w, r)
				}
				if iError.Error() == sdkerror.ErrConsentRequired.Name && iError.GetReason() == "" {
					d.PageResponseHandler.DisplayConsentPage(w, r)
				}
				authRequestContext.SetError(iError)
				err := d.ErrorWriter.WriteRedirectError(authRequestContext, w, r)
				if err != nil {
					d.ErrorStrategy(err, w)
				}
				return
			}
		}
		if err := d.ResponseWriter.WriteAuthorizationResponse(authRequestContext, w, r); err != nil {
			d.ErrorStrategy(err, w)
		}
	}
}
