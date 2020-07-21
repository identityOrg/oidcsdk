package manager

import "net/http"

func (d *DefaultManager) ProcessIntrospectionEP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	if introspectionRequestContext, iError := d.IntrospectionRequestContextFactory(request); iError != nil {
		err := d.JsonErrorWriter(iError, writer, request)
		if err != nil {
			d.ErrorStrategy(err, writer)
		}
		return
	} else {
		for _, handler := range d.IntrospectionEPHandlers {
			if iError := handler.HandleIntrospectionEP(ctx, introspectionRequestContext); iError != nil {
				err := d.JsonErrorWriter(iError, writer, request)
				if err != nil {
					d.ErrorStrategy(err, writer)
				}
				return
			}
		}
		if err := d.IntrospectionResponseWriter(introspectionRequestContext, writer, request); err != nil {
			d.ErrorStrategy(err, writer)
		}
	}
}
