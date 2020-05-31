package oauth2_oidc_sdk

import (
	"net/url"
)

type AuthenticationRequest struct {
	AuthorizationRequest
	ResponseMode ResponseModeType
	Nonce        string
	Display      DisplayType
	MaxAge       int
	IdTokenHint  string
	LoginHint    string
	AcrValues    string
	Purpose      string
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
