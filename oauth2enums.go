package oauth2_oidc_sdk

import "strings"

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
	var buf strings.Builder
	for _, id := range values {
		if buf.Len() > 0 {
			buf.WriteString(id.StaringValue())
		} else {
			buf.WriteString(" ")
			buf.WriteString(id.StaringValue())
		}
	}
	return buf.String()
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

type PromptType string

func (st PromptType) StaringValue() string {
	return string(st)
}

const (
	PromptNone          PromptType = "none"
	PromptConsent       PromptType = "consent"
	PromptLogin         PromptType = "login"
	PromptSelectAccount PromptType = "select_account"
)

type PromptTypeArray []PromptType

func (values PromptTypeArray) StaringValue() string {
	var buf strings.Builder
	for _, id := range values {
		if buf.Len() > 0 {
			buf.WriteString(id.StaringValue())
		} else {
			buf.WriteString(" ")
			buf.WriteString(id.StaringValue())
		}
	}
	return buf.String()
}

type CodeChallengeMethodType string

func (rm CodeChallengeMethodType) StaringValue() string {
	return string(rm)
}

const (
	CodeChallengeMethodPlain CodeChallengeMethodType = "plain"
	CodeChallengeMethodS256  CodeChallengeMethodType = "S256"
)
