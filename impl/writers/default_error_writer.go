package writers

import (
	"encoding/json"
	"errors"
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
	"net/url"
)

type DefaultErrorWriter struct {
}

func NewDefaultErrorWriter() *DefaultErrorWriter {
	return &DefaultErrorWriter{}
}

func (*DefaultErrorWriter) WriteJsonError(pError sdk.IError, additionalValues url.Values, writer http.ResponseWriter, r *http.Request) error {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeJson)
	writer.WriteHeader(pError.GetStatusCode())

	err := json.NewEncoder(writer).Encode(pError)
	return err
}

func (*DefaultErrorWriter) WriteRedirectError(requestContext sdk.IAuthenticationRequestContext, writer http.ResponseWriter, request *http.Request) error {
	mode := requestContext.GetResponseMode()
	switch mode {
	case sdk.ResponseModeFragment:
		form := buildErrorResponseForm(requestContext)
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.Fragment = form.Encode()
		http.Redirect(writer, request, redirectUri.String(), http.StatusFound)
		return nil
	case sdk.ResponseModeQuery:
		form := buildErrorResponseForm(requestContext)
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.RawQuery = form.Encode()
		http.Redirect(writer, request, redirectUri.String(), http.StatusFound)
		return nil
	}
	return errors.New("invalid response mode")
}

func (*DefaultErrorWriter) WriteBearerError(pError sdk.IError, additionalValues url.Values, writer http.ResponseWriter, request *http.Request) error {
	panic("implement me")
}

func buildErrorResponseForm(requestContext sdk.IAuthenticationRequestContext) url.Values {
	form := url.Values{}
	err := requestContext.GetError()
	form.Add("error", err.Error())
	form.Add("error_description", err.GetDescription())
	if requestContext.GetState() != "" {
		form.Add("state", requestContext.GetState())
	}
	return form
}
