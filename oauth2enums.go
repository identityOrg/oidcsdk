package oauth2_oidc_sdk

import (
	"net/url"
	"reflect"
	"strings"
)

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
	String() string
}

type ResponseTypeArray []string

func (values ResponseTypeArray) String() string {
	return strings.Join(values, " ")
}

const (
	ResponseTypeCode    = "code"
	ResponseTypeToken   = "token"
	ResponseTypeIdToken = "id_token"
)

const (
	ScopeTypeOpenId        = "openid"
	ScopeTypeProfile       = "profile"
	ScopeTypeOfflineAccess = "offline_access"
)

type ScopeTypeArray []string

func (values ScopeTypeArray) String() string {
	return strings.Join(values, " ")
}

const (
	GTAuthorizationCode     = "authorization_code"
	GTImplicit              = "implicit"
	GTResourceOwnerPassword = "password"
	GTClientCredential      = "client_credentials"
)

const (
	PromptNone          = "none"
	PromptConsent       = "consent"
	PromptLogin         = "login"
	PromptSelectAccount = "select_account"
)

type PromptTypeArray []string

func (values PromptTypeArray) String() string {
	return strings.Join(values, " ")
}

func SpacesStringArrayEncoder(value reflect.Value) string {
	var strValues []string
	if reflect.Slice == value.Kind() {
		for i := 0; i < value.Len(); i++ {
			strValues = append(strValues, value.Index(i).String())
		}
	}
	return strings.Join(strValues, " ")
}

func SpacesStringArrayDecoder(s string) reflect.Value {
	split := strings.Split(s, " ")
	value := make([]string, 0)
	for _, v := range split {
		value = append(value, v)
	}
	return reflect.ValueOf(value)
}

const (
	CodeChallengeMethodPlain = "plain"
	CodeChallengeMethodS256  = "S256"
)

type UrlType url.URL

func (ut UrlType) String() string {
	var ulUt url.URL
	ulUt = url.URL(ut)
	return ulUt.String()
}

func UrlEncoder(value reflect.Value) string {
	strValue := value.MethodByName("String").Call([]reflect.Value{})
	return strValue[0].String()
}

func UrlDecoder(s string) reflect.Value {
	var z reflect.Value
	parse, err := url.Parse(s)
	if err != nil {
		return z
	}
	return reflect.ValueOf(UrlType(*parse))
}
