package manager

import "net/http"

func (d *DefaultManager) ProcessIntrospectionEP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	if requestContext, iError := d.IntrospectionRequestContextFactory(request); iError != nil {
		err := d.JsonErrorWriter(iError, writer, request)
		if err != nil {
			d.ErrorStrategy(err, writer)
		}
		return
	} else {
		for _, handler := range d.IntrospectionEPHandlers {
			if iError := handler.HandleIntrospectionEP(ctx, requestContext); iError != nil {
				err := d.JsonErrorWriter(iError, writer, request)
				if err != nil {
					d.ErrorStrategy(err, writer)
				}
				return
			}
		}
		if err := d.IntrospectionResponseWriter(requestContext, writer, request); err != nil {
			d.ErrorStrategy(err, writer)
		}
	}
}
