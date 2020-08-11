package manager

import "net/http"

func (d *DefaultManager) ProcessUserInfoEP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	if requestContext, iError := d.RequestContextFactory.BuildUserInfoRequestContext(request); iError != nil {
		err := d.ErrorWriter.WriteJsonError(iError, nil, writer, request)
		if err != nil {
			d.ErrorStrategy(err, writer)
		}
		return
	} else {
		for _, handler := range d.UserInfoEPHandlers {
			if iError := handler.HandleUserInfoEP(ctx, requestContext); iError != nil {
				err := d.ErrorWriter.WriteJsonError(iError, nil, writer, request)
				if err != nil {
					d.ErrorStrategy(err, writer)
				}
				return
			}
		}
		if err := d.ResponseWriter.WriteUserInfoResponse(requestContext, writer, request); err != nil {
			d.ErrorStrategy(err, writer)
		}
	}
}
