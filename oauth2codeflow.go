package oauth2_oidc_sdk

import (
	"net/url"
)

type (
	AuthorizationRequest struct {
		RequestUri          UrlType           `schema:"-"`
		ClientId            string            `schema:"client_id,required"`
		Scopes              ScopeTypeArray    `schema:"scope"`
		ResponseType        ResponseTypeArray `schema:"response_type,required"`
		State               string            `schema:"state"`
		RedirectUri         UrlType           `schema:"redirect_uri"`
		CodeChallenge       string            `schema:"code_challenge"`
		CodeChallengeMethod string            `schema:"code_challenge_method"`
		Prompt              PromptTypeArray   `schema:"prompt"`
	}
	AuthorizationResponse struct {
	}
)

func (ar AuthorizationRequest) Render() url.URL {
	renderUrl := ar.RequestUri
	queryValue := url.Values{}
	queryValue.Add(ParameterClientId, ar.ClientId)
	queryValue.Add(ParameterState, ar.State)
	queryValue.Add(ParameterResponseType, ar.ResponseType.String())
	queryValue.Add(ParameterScope, ar.Scopes.String())
	queryValue.Add(ParameterRedirectUri, ar.RedirectUri.String())
	queryValue.Add(ParameterCodeChallenge, ar.CodeChallenge)
	queryValue.Add(ParameterCodeChallengeMethod, ar.CodeChallengeMethod)
	queryValue.Add(ParameterPrompt, ar.Prompt.String())
	renderUrl.RawQuery = queryValue.Encode()
	return url.URL(renderUrl)
}

func (ar AuthorizationRequest) InferGrantType() string {
	for _, rt := range ar.ResponseType {
		if rt == ResponseTypeCode {
			return GTAuthorizationCode
		}
	}
	return GTImplicit
}

//InferResponseMode is a method to determine response mode
func (ar AuthorizationRequest) InferResponseMode() ResponseModeType {
	for _, rt := range ar.ResponseType {
		if rt == ResponseTypeCode {
			return ResponseModeQuery
		}
	}
	return ResponseModeFragment
}

func ParseAuthorizationRequest(reqUrl url.URL) (AuthorizationRequest, error) {
	//query := reqUrl.Query()
	request := AuthorizationRequest{}
	return request, nil
}
