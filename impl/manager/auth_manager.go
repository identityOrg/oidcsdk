package manager

import (
	"net/http"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

func (d *DefaultManager) ProcessAuthorizationEP(w http.ResponseWriter, r *http.Request) sdk.Result {
	ctx := r.Context()
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
		return sdk.ResultNoOperation
	} else {
		if sess, err := d.UserSessionManager.RetrieveUserSession(r); err == nil {
			authRequestContext.SetUserSession(sess)
		}
		for _, handler := range d.AuthEPHandlers {
			if iError := handler.HandleAuthEP(ctx, authRequestContext); iError != nil {
				if iError.Error() == sdkerror.ErrLoginRequired.Name && iError.GetReason() == "" {
					return sdk.ResultLoginRequired
				}
				if iError.Error() == sdkerror.ErrConsentRequired.Name && iError.GetReason() == "" {
					return sdk.ResultConsentRequired
				}
				authRequestContext.SetError(iError)
				err := d.AuthenticationErrorWriter(authRequestContext, w, r)
				if err != nil {
					d.ErrorStrategy(err, w)
				}
				return sdk.ResultNoOperation
			}
		}
		if err := d.AuthenticationResponseWriter(authRequestContext, w, r); err != nil {
			d.ErrorStrategy(err, w)
		}
		return sdk.ResultNoOperation
	}
}
