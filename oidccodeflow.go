package oauth2_oidc_sdk

import (
	"net/url"
)

type AuthenticationRequest struct {
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
type AuthenticationResponse struct {
	AuthorizationResponse
}

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
