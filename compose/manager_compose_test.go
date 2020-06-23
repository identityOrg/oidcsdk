package compose

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/memdbstore"
	"oauth2-oidc-sdk/impl/strategies"
	"oauth2-oidc-sdk/util"
	"strings"
	"testing"
)

func TestDefaultManager(t *testing.T) {
	config := sdk.NewConfig("http://localhost:8080")
	private, public := util.GenerateRSAKeyPair()
	strategy := strategies.NewDefaultStrategy(private, public)
	sequence := CreateDefaultSequence()
	sequence = append(sequence, memdbstore.NewInMemoryDB(true))
	got := DefaultManager(config, strategy, sequence...)
	rw := httptest.NewRecorder()
	form := url.Values{}
	form.Set("grant_type", "password")
	form.Set("username", "user")
	form.Set("password", "user")
	request := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	got.ProcessTokenEP(rw, request)
	println(rw.Code)
}
