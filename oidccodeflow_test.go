package oauth2_oidc_sdk

import (
	"fmt"
	"net/url"
	"testing"
)

func TestOIdcCode(t *testing.T) {
	ar := AuthenticationRequest{}
	reqUrl, _ := url.Parse("http://localhost:8080")
	ar.RequestUri = UrlType(*reqUrl)
	ar.ClientId = RandomIdString(20)
	ar.ResponseType = ResponseTypeArray{ResponseTypeCode}
	ar.State = RandomIdString(10)
	ar.Scopes = ScopeTypeArray{ScopeTypeOpenId, ScopeTypeProfile, ScopeTypeOfflineAccess}
	ar.Prompt = PromptTypeArray{PromptConsent, PromptLogin, PromptSelectAccount}
	ar.ResponseMode = ResponseModeQuery
	ar.CodeChallengeMethod = CodeChallengeMethodS256
	ar.CodeChallenge = RandomIdString(16)
	ar.Display = DisplayPage
	ar.MaxAge = 120
	render := ar.Render()
	fmt.Println(render.String())
}
