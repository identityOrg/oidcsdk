package oauth2_oidc_sdk

import "github.com/gorilla/schema"

func createFormEncoder() *schema.Encoder {
	encoder := schema.NewEncoder()
	encoder.RegisterEncoder(UrlType{}, UrlEncoder)
	encoder.RegisterEncoder(PromptTypeArray{}, SpacesStringArrayEncoder)
	encoder.RegisterEncoder(ScopeTypeArray{}, SpacesStringArrayEncoder)
	return encoder
}

func createFormDecoder() *schema.Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.RegisterConverter(UrlType{}, UrlDecoder)
	decoder.RegisterConverter(PromptTypeArray{}, SpacesStringArrayDecoder)
	decoder.RegisterConverter(ScopeTypeArray{}, SpacesStringArrayDecoder)
	return decoder
}
