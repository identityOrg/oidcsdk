package oauth2_oidc_sdk

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestAuthorizationRequest_Render(t *testing.T) {
	ar := AuthorizationRequest{}
	reqUrl, _ := url.Parse("http://localhost:8080")
	ar.RequestUri = UrlType(*reqUrl)
	ar.ClientId = RandomIdString(20)
	ar.ResponseType = ResponseTypeArray{ResponseTypeCode}
	ar.State = RandomIdString(10)
	ar.Scopes = ScopeTypeArray{ScopeTypeOpenId, ScopeTypeProfile, ScopeTypeOfflineAccess}
	render, _ := ar.Render()
	state := render.Query().Get("state")
	assert.Equal(t, state, ar.State)
	rt := render.Query().Get("response_type")
	assert.Equal(t, rt, ResponseTypeCode)
}

func TestAuthorizationRequest_Parse(t *testing.T) {
	testUrl, _ := url.Parse("http://localhost:8080?client_id=8tjpe883tv1p0s3gbngo&code_challenge=&code_challenge_method=&prompt=login&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fredirect&response_type=token&scope=openid+profile+offline_access&state=mptwwwp0fa")
	ar := AuthorizationRequest{}
	err := ar.Parse(*testUrl)
	if err != nil {
		t.Fail()
		return
	}
	assert.Equal(t, ar.RequestUri.String(), "http://localhost:8080")
	assert.Equal(t, ar.Prompt, PromptTypeArray{PromptLogin})
}

func TestAuthorizationSuccessResponse_Parse(t *testing.T) {
	testUrl, _ := url.Parse("http://localhost:8080/redirect#access_token=dms6hg26hj&authorization_code=4kungm3qc4yqbyfn5oae&state=21xr4opsq0")
	asr := AuthorizationSuccessResponse{}
	err := asr.Parse(*testUrl)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, asr.ResponseMode.String(), ResponseModeFragment.String())
}

func TestAuthorizationSuccessResponse_Render(t *testing.T) {
	redUri, _ := url.Parse("http://localhost:8080/redirect")
	asr := AuthorizationSuccessResponse{
		RedirectUri:       UrlType(*redUri),
		State:             RandomIdString(10),
		ResponseMode:      ResponseModeQuery,
		AuthorizationCode: RandomIdString(20),
		AccessToken:       RandomIdString(10),
	}

	render, err := asr.Render()
	if err != nil {
		t.Fatal("render failed")
	}

	assert.Greater(t, len(render.RawQuery), 0)
	assert.Len(t, render.Fragment, 0)

	values := render.Query()

	assert.NotNil(t, values.Get("access_token"))
}

func TestAuthorizationSuccessResponse_Render2(t *testing.T) {
	redUri, _ := url.Parse("http://localhost:8080/redirect")
	asr := AuthorizationSuccessResponse{
		RedirectUri:       UrlType(*redUri),
		State:             RandomIdString(10),
		ResponseMode:      ResponseModeFragment,
		AuthorizationCode: RandomIdString(20),
		AccessToken:       RandomIdString(10),
	}

	render, err := asr.Render()
	if err != nil {
		t.Fatal("render failed")
	}

	assert.Greater(t, len(render.Fragment), 0)
	assert.Len(t, render.RawQuery, 0)
}
