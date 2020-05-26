package oauth2_oidc_sdk

import (
	"net/url"
	"strconv"
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

func (ar AuthenticationRequest) Render() url.URL {
	renderUrl := ar.AuthorizationRequest.Render()
	queryValue := renderUrl.Query()
	queryValue.Add(ParameterNonce, ar.Nonce)
	queryValue.Add(ParameterDisplay, ar.Display.StaringValue())
	queryValue.Add(ParameterMaxAge, strconv.Itoa(ar.MaxAge))
	queryValue.Add(ParameterIdTokenHint, ar.IdTokenHint)
	queryValue.Add(ParameterLoginHint, ar.LoginHint)
	queryValue.Add(ParameterAcrValues, ar.AcrValues)
	queryValue.Add(ParameterPurpose, ar.Purpose)
	renderUrl.RawQuery = queryValue.Encode()
	return renderUrl
}

func (ar AuthenticationRequest) InferResponseMode() ResponseModeType {
	if ar.ResponseMode != "" {
		return ar.ResponseMode
	} else {
		return ar.AuthorizationRequest.InferResponseMode()
	}
}
