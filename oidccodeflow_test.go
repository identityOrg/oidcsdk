package oauth2_oidc_sdk

import (
	"github.com/magiconair/properties/assert"
	"net/url"
	"testing"
)

func TestAuthenticationRequest_Render(t *testing.T) {
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
	render, _ := ar.Render()
	state := render.Query().Get("state")
	assert.Equal(t, state, ar.State)
}

func TestAuthenticationRequest_Parse(t *testing.T) {
	testUrl, _ := url.Parse("http://localhost:8080?client_id=8tjpe883tv1p0s3gbngo&code_challenge=&code_challenge_method=&prompt=login&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fredirect&response_type=token&scope=openid+profile+offline_access&state=mptwwwp0fa&display=popup")
	ar := AuthenticationRequest{}
	err := ar.Parse(*testUrl)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, ar.RequestUri.String(), "http://localhost:8080")
	assert.Equal(t, ar.Display, DisplayPopup)
}
