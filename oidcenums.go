package oauth2_oidc_sdk

const (
	ParameterNonce       = "nonce"
	ParameterDisplay     = "display"
	ParameterMaxAge      = "max_age"
	ParameterIdTokenHint = "id_token_hint"
	ParameterLoginHint   = "login_hint"
	ParameterAcrValues   = "acr_values"
	ParameterPurpose     = "purpose"
)

type ResponseModeType string

func (rm ResponseModeType) String() string {
	return string(rm)
}

const (
	ResponseModeQuery    ResponseModeType = "query"
	ResponseModeFragment ResponseModeType = "fragment"
	ResponseModePost     ResponseModeType = "post"
)

type DisplayType string

func (rm DisplayType) String() string {
	return string(rm)
}

const (
	DisplayPage  DisplayType = "page"
	DisplayPopup DisplayType = "popup"
	DisplayTouch DisplayType = "touch"
	DisplayWap   DisplayType = "wap"
)
