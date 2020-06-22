package compose

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/tokens"
	"oauth2-oidc-sdk/util"
	"strings"
	"testing"
)

func TestDefaultManager(t *testing.T) {
	config := sdk.Config{}
	private, public := util.GenerateRSAKeyPair()
	strategy := tokens.NewDefaultStrategy(private, public)
	got := DefaultManager(config, strategy, nil)
	rw := httptest.NewRecorder()
	form := url.Values{}
	form.Set("grant_type", "password")
	form.Set("username", "user")
	form.Set("password", "user")
	request := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	got.ProcessTokenEP(rw, request)
}
