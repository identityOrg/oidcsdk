package manager

import "net/http"

func (d *DefaultManager) ProcessRevocationEP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	if requestContext, iError := d.RequestContextFactory.BuildRevocationRequestContext(request); iError != nil {
		err := d.ErrorWriter.WriteJsonError(iError, nil, writer, request)
		if err != nil {
			d.ErrorStrategy(err, writer)
		}
		return
	} else {
		for _, handler := range d.RevocationEPHandlers {
			if iError := handler.HandleRevocationEP(ctx, requestContext); iError != nil {
				err := d.ErrorWriter.WriteJsonError(iError, nil, writer, request)
				if err != nil {
					d.ErrorStrategy(err, writer)
				}
				return
			}
		}
		if err := d.ResponseWriter.WriteRevocationResponse(writer, request); err != nil {
			d.ErrorStrategy(err, writer)
		}
	}
}
