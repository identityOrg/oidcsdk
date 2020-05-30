package oauth2_oidc_sdk

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/url"
	"testing"
)

func TestCode(t *testing.T) {
	ar := AuthorizationRequest{}
	reqUrl, _ := url.Parse("http://localhost:8080")
	ar.RequestUri = UrlType(*reqUrl)
	ar.ClientId = RandomIdString(20)
	ar.ResponseType = ResponseTypeArray{ResponseTypeCode}
	ar.State = RandomIdString(10)
	ar.Scopes = ScopeTypeArray{ScopeTypeOpenId, ScopeTypeProfile, ScopeTypeOfflineAccess}
	render := ar.Render()
	fmt.Println(render.String())
}

func TestParseAuthorizationRequest(t *testing.T) {
	ar := AuthorizationRequest{}
	reqUrl, _ := url.Parse("http://localhost:8080")
	rediUri, _ := url.Parse("http://localhost:8080/redirect")
	ar.RequestUri = UrlType(*reqUrl)
	ar.RedirectUri = UrlType(*rediUri)
	ar.ClientId = RandomIdString(20)
	ar.ResponseType = ResponseTypeArray{ResponseTypeToken}
	ar.State = RandomIdString(10)
	ar.Scopes = ScopeTypeArray{ScopeTypeOpenId, ScopeTypeProfile, ScopeTypeOfflineAccess}
	ar.Prompt = PromptTypeArray{PromptLogin}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.RegisterConverter(UrlType{}, UrlDecoder)
	decoder.RegisterConverter(PromptTypeArray{}, SpacesStringArrayDecoder)
	decoder.RegisterConverter(ScopeTypeArray{}, SpacesStringArrayDecoder)
	encoder := schema.NewEncoder()
	encoder.RegisterEncoder(UrlType{}, UrlEncoder)
	encoder.RegisterEncoder(PromptTypeArray{}, SpacesStringArrayEncoder)
	encoder.RegisterEncoder(ScopeTypeArray{}, SpacesStringArrayEncoder)

	query := url.Values{}
	err := encoder.Encode(ar, query)

	if err != nil {
		t.Fatal(err)
	}

	ar1 := AuthorizationRequest{}
	err = decoder.Decode(&ar1, query)

	if err != nil {
		t.Fatal(err)
	}

	renderUrl := ar1.Render()
	println(renderUrl.String())
}
