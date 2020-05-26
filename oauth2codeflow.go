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
	CodeChallengeMethod CodeChallengeMethodType
	Prompt              PromptTypeArray
}
type AuthorizationResponse struct {
}

func (ar AuthorizationRequest) Render() url.URL {
	renderUrl := ar.RequestUri
	queryValue := url.Values{}
	queryValue.Add(ParameterClientId, ar.ClientId)
	queryValue.Add(ParameterState, ar.State)
	queryValue.Add(ParameterResponseType, ar.ResponseType.StaringValue())
	queryValue.Add(ParameterScope, ar.Scopes.StaringValue())
	queryValue.Add(ParameterRedirectUri, ar.RedirectUri.String())
	queryValue.Add(ParameterCodeChallenge, ar.CodeChallenge)
	queryValue.Add(ParameterCodeChallengeMethod, ar.CodeChallengeMethod.StaringValue())
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

func (ar AuthorizationRequest) InferResponseMode() ResponseModeType {
	switch ar.ResponseType {
	case ResponseTypeCode:
		return ResponseModeQuery
	default:
		return ResponseModeFragment
	}
}
