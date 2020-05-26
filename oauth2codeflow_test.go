package oauth2_oidc_sdk

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCode(t *testing.T) {
	ar := AuthorizationRequest{}
	reqUrl, _ := url.Parse("http://localhost:8080")
	ar.RequestUri = *reqUrl
	ar.ClientId = RandomIdString(20)
	ar.ResponseType = ResponseTypeIdTokenToken
	ar.State = RandomIdString(10)
	ar.Scopes = ScopeTypeArray{ScopeTypeOpenId, ScopeTypeProfile, ScopeTypeOfflineAccess}
	render := ar.Render()
	fmt.Println(render.String())
}
