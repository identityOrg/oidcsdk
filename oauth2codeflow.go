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

func (ar AuthorizationRequest) Render() (*url.URL, error) {
	var renderUrl = url.URL(ar.RequestUri)
	values := url.Values{}
	err := createFormEncoder().Encode(ar, values)
	if err != nil {
		return nil, err
	}
	renderUrl.RawQuery = values.Encode()
	return &renderUrl, nil
}

func (ar *AuthorizationRequest) Parse(reqUrl url.URL) error {
	query := reqUrl.Query()
	reqUrl.RawQuery = ""
	ar.RequestUri = UrlType(reqUrl)
	return createFormDecoder().Decode(ar, query)
}
