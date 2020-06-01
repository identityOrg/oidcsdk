package oauth2_oidc_sdk

import (
	"github.com/stretchr/testify/assert"
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
	testUrl, _ := url.Parse("http://localhost:8080?acr_values=&client_id=i8paocx7nzyy0cici9e0&code_challenge=stppsjnszj8bdi26&code_challenge_method=S256&display=page&id_token_hint=&login_hint=&max_age=120&nonce=&prompt=consent+login+select_account&purpose=&redirect_uri=&response_mode=query&response_type=code&scope=openid+profile+offline_access&state=fj1aihb076&display=popup")
	ar := AuthenticationRequest{}
	err := ar.Parse(*testUrl)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, ar.RequestUri.String(), "http://localhost:8080")
	assert.Equal(t, ar.Display, DisplayPopup)
}
