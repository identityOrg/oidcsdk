package manager

import "net/http"

func (d *DefaultManager) ProcessIntrospectionEP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	if requestContext, iError := d.RequestContextFactory.BuildIntrospectionRequestContext(request); iError != nil {
		err := d.ErrorWriter.WriteJsonError(iError, nil, writer, request)
		if err != nil {
			d.ErrorStrategy(err, writer)
		}
		return
	} else {
		for _, handler := range d.IntrospectionEPHandlers {
			if iError := handler.HandleIntrospectionEP(ctx, requestContext); iError != nil {
				err := d.ErrorWriter.WriteJsonError(iError, nil, writer, request)
				if err != nil {
					d.ErrorStrategy(err, writer)
				}
				return
			}
		}
		if err := d.ResponseWriter.WriteIntrospectionResponse(requestContext, writer, request); err != nil {
			d.ErrorStrategy(err, writer)
		}
	}
}
