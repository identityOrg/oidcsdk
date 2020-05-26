package oauth2_oidc_sdk

const (
	ParameterClientId            = "client_id"
	ParameterState               = "state"
	ParameterScope               = "scope"
	ParameterResponseType        = "response_type"
	ParameterRedirectUri         = "redirect_uri"
	ParameterCodeChallenge       = "code_challenge"
	ParameterCodeChallengeMethod = "code_challenge_method"
	ParameterResource            = "resource"
	ParameterPrompt              = "prompt"
)

type Id interface {
	StaringValue() string
}

type ResponseType string

func (rt ResponseType) StaringValue() string {
	return string(rt)
}

const (
	ResponseTypeCode             ResponseType = "code"
	ResponseTypeToken            ResponseType = "token"
	ResponseTypeIdToken          ResponseType = "id_token"
	ResponseTypeIdTokenToken     ResponseType = "id_token token"
	ResponseTypeIdTokenTokenCode ResponseType = "id_token token code"
)

type ScopeType string

func (st ScopeType) StaringValue() string {
	return string(st)
}

const (
	ScopeTypeOpenId        ScopeType = "openid"
	ScopeTypeProfile       ScopeType = "profile"
	ScopeTypeOfflineAccess ScopeType = "offline_access"
)

type ScopeTypeArray []ScopeType

func (values ScopeTypeArray) StaringValue() string {
	var finalValue string
	for _, id := range values {
		if len(finalValue) > 0 {
			finalValue = finalValue + " " + id.StaringValue()
		} else {
			finalValue = id.StaringValue()
		}
	}
	return finalValue
}

type GrantType string

func (gt GrantType) StaringValue() string {
	return string(gt)
}

const (
	GTAuthorizationCode     GrantType = "authorization_code"
	GTImplicit              GrantType = "implicit"
	GTResourceOwnerPassword GrantType = "password"
	GTClientCredential      GrantType = "client_credentials"
)

type ResponseMode string

func (rm ResponseMode) StaringValue() string {
	return string(rm)
}

const (
	ResponseModeQuery    ResponseMode = "query"
	ResponseModeFragment ResponseMode = "fragment"
	ResponseModePost     ResponseMode = "post"
)

type PromptType string

func (st PromptType) StaringValue() string {
	return string(st)
}

const (
	PromptTypeNone          PromptType = "none"
	PromptTypeConsent       PromptType = "consent"
	PromptTypeLogin         PromptType = "login"
	PromptTypeSelectAccount PromptType = "select_account"
)

type PromptTypeArray []PromptType

func (values PromptTypeArray) StaringValue() string {
	var finalValue string
	for _, id := range values {
		if len(finalValue) > 0 {
			finalValue = finalValue + " " + id.StaringValue()
		} else {
			finalValue = id.StaringValue()
		}
	}
	return finalValue
}
