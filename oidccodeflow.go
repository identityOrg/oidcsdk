package oauth2_oidc_sdk

import (
	"net/url"
)

type (
	AuthenticationRequest struct {
		AuthorizationRequest
		ResponseMode ResponseModeType `schema:"response_mode"`
		Nonce        string           `schema:"nonce"`
		Display      DisplayType      `schema:"display"`
		MaxAge       int              `schema:"max_age"`
		IdTokenHint  string           `schema:"id_token_hint"`
		LoginHint    string           `schema:"login_hint"`
		AcrValues    string           `schema:"acr_values"`
		Purpose      string           `schema:"purpose"`
	}
	AuthenticationSuccessResponse struct {
		AuthorizationSuccessResponse
		IdToken      string `schema:"id_token"`
		SessionState string `schema:"session_state"`
	}
	AuthenticationErrorResponse struct {
		AuthorizationErrorResponse
	}
)

func (ar AuthenticationRequest) InferResponseMode() ResponseModeType {
	if ar.ResponseMode != "" {
		return ar.ResponseMode
	} else {
		return ar.AuthorizationRequest.InferResponseMode()
	}
}

func (ar AuthenticationRequest) Render() (*url.URL, error) {
	var renderUrl = url.URL(ar.RequestUri)
	values := url.Values{}
	err := createFormEncoder().Encode(ar, values)
	if err != nil {
		return nil, err
	}
	renderUrl.RawQuery = values.Encode()
	return &renderUrl, nil
}

func (ar *AuthenticationRequest) Parse(reqUrl url.URL) error {
	query := reqUrl.Query()
	reqUrl.RawQuery = ""
	ar.RequestUri = UrlType(reqUrl)
	return createFormDecoder().Decode(ar, query)
}

func (asr AuthenticationSuccessResponse) Render() (*url.URL, error) {
	var renderUrl = url.URL(asr.RedirectUri)
	values := url.Values{}
	err := createFormEncoder().Encode(asr, values)
	if err != nil {
		return nil, err
	}
	if asr.ResponseMode == ResponseModeFragment {
		renderUrl.Fragment = values.Encode()
	} else {
		renderUrl.RawQuery = values.Encode()
	}
	return &renderUrl, nil
}

func (asr *AuthenticationSuccessResponse) Parse(reqUrl url.URL) error {
	query := reqUrl.Query()
	reqUrl.RawQuery = ""
	asr.RedirectUri = UrlType(reqUrl)
	return createFormDecoder().Decode(asr, query)
}

func (aer AuthenticationErrorResponse) Render() (*url.URL, error) {
	var renderUrl = url.URL(aer.RedirectUri)
	values := url.Values{}
	err := createFormEncoder().Encode(aer, values)
	if err != nil {
		return nil, err
	}
	if aer.ResponseMode == ResponseModeFragment {
		renderUrl.Fragment = values.Encode()
	} else {
		renderUrl.RawQuery = values.Encode()
	}
	return &renderUrl, nil
}

func (aer *AuthenticationErrorResponse) Parse(reqUrl url.URL) error {
	query := reqUrl.Query()
	reqUrl.RawQuery = ""
	aer.RedirectUri = UrlType(reqUrl)
	return createFormDecoder().Decode(aer, query)
}
