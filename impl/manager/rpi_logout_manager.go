package manager

import (
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"net/http"
)

func (d *DefaultManager) ProcessRPILogoutEP(w http.ResponseWriter, r *http.Request) {
	if requestContext, iError := d.RequestContextFactory.BuildRPILogoutRequestContext(r); iError != nil {
		d.PageResponseHandler.DisplayErrorPage(iError, w, r)
	} else {
		if sess, err := d.UserSessionManager.RetrieveUserSession(r); err == nil {
			requestContext.SetUserSession(sess)
		}
		ctx := r.Context()
		for _, handler := range d.RPILogoutEPHandlers {
			if iError := handler.HandleRPILogoutEP(ctx, requestContext); iError != nil {
				if iError.Error() == sdkerror.ErrInteractionRequired.Name && iError.GetReason() == "" {
					d.PageResponseHandler.DisplayLogoutConsentPage(w, r)
				}
				d.PageResponseHandler.DisplayErrorPage(iError, w, r)
				return
			}
		}
		//todo actual logout
		if requestContext.GetPostLogoutRedirectUri() != "" {
			d.ResponseWriter.WriteRPILogoutResponse(requestContext, w, r)
		} else {
			d.PageResponseHandler.DisplayLogoutStatusPage(w, r)
		}
	}
}
