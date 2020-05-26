package oauth2_oidc_sdk

import (
	"net/url"
)

type AuthorizationRequest struct {
	RequestUri          url.URL
	ClientId            string
	Scopes              ScopeTypeArray
	ResponseType        ResponseType
	State               string
	RedirectUri         url.URL
	CodeChallenge       string
	CodeChallengeMethod string
	Prompt              PromptTypeArray
}
type AuthorizationResponse struct {
}
type AuthenticationRequest struct {
	AuthorizationRequest
	ResponseMode ResponseMode
}
type AuthenticationResponse struct {
	AuthorizationResponse
}

func (ar AuthorizationRequest) Render() (*url.URL, error) {
	return url.Parse("http://localhost:8080/authorize")
}

func (ar AuthenticationRequest) Render() url.URL {
	renderUrl := ar.RequestUri
	queryValue := url.Values{}
	queryValue.Add(ParameterClientId, ar.ClientId)
	queryValue.Add(ParameterState, ar.State)
	queryValue.Add(ParameterResponseType, ar.ResponseType.StaringValue())
	queryValue.Add(ParameterScope, ar.Scopes.StaringValue())
	queryValue.Add(ParameterRedirectUri, ar.RedirectUri.String())
	queryValue.Add(ParameterCodeChallenge, ar.CodeChallenge)
	queryValue.Add(ParameterCodeChallengeMethod, ar.CodeChallengeMethod)
	queryValue.Add(ParameterPrompt, ar.Prompt.StaringValue())
	renderUrl.RawQuery = queryValue.Encode()
	return renderUrl
}

func (ar AuthorizationRequest) InferGrantType() GrantType {
	switch ar.ResponseType {
	case ResponseTypeCode:
		return GTAuthorizationCode
	default:
		return GTImplicit
	}
}

func (ar AuthorizationRequest) InferResponseMode() ResponseMode {
	switch ar.ResponseType {
	case ResponseTypeCode:
		return ResponseModeQuery
	default:
		return ResponseModeFragment
	}
}

func (ar AuthenticationRequest) InferResponseMode() ResponseMode {
	if ar.ResponseMode != "" {
		return ar.ResponseMode
	} else {
		return ar.AuthorizationRequest.InferResponseMode()
	}
}
