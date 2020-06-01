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
	AuthorizationSuccessResponse struct {
		RedirectUri       UrlType          `schema:"-"`
		State             string           `schema:"state"`
		ResponseMode      ResponseModeType `schema:"-"`
		AuthorizationCode string           `schema:"authorization_code"`
		AccessToken       string           `schema:"access_token"`
	}
	AuthorizationErrorResponse struct {
		RedirectUri    UrlType          `schema:"-"`
		State          string           `schema:"state"`
		ResponseMode   ResponseModeType `schema:"-"`
		ErrorCode      string           `schema:"code"`
		Description    string           `schema:"description"`
		HttpStatusCode int              `schema:"-"`
		Url            UrlType          `schema:"uri"`
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

func (asr AuthorizationSuccessResponse) Render() (*url.URL, error) {
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

func (asr *AuthorizationSuccessResponse) Parse(reqUrl url.URL) error {
	var query url.Values
	if reqUrl.RawQuery != "" {
		asr.ResponseMode = ResponseModeQuery
		query = reqUrl.Query()
	} else if reqUrl.Fragment != "" {
		asr.ResponseMode = ResponseModeFragment
		reqUrl.RawQuery = reqUrl.Fragment
		query = reqUrl.Query()
		reqUrl.Fragment = ""
	}
	reqUrl.RawQuery = ""
	asr.RedirectUri = UrlType(reqUrl)
	return createFormDecoder().Decode(asr, query)
}

func (aer AuthorizationErrorResponse) Render() (*url.URL, error) {
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

func (aer *AuthorizationErrorResponse) Parse(reqUrl url.URL) error {
	var query url.Values
	if reqUrl.RawQuery != "" {
		aer.ResponseMode = ResponseModeQuery
		query = reqUrl.Query()
	} else if reqUrl.Fragment != "" {
		aer.ResponseMode = ResponseModeFragment
		reqUrl.RawQuery = reqUrl.Fragment
		query = reqUrl.Query()
		reqUrl.Fragment = ""
	}
	reqUrl.RawQuery = ""
	aer.RedirectUri = UrlType(reqUrl)
	return createFormDecoder().Decode(aer, query)
}
