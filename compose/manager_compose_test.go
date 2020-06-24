package compose

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
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

func TestDefaultManager_Password(t *testing.T) {
	got := CreateManager()
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

func TestDefaultManager_Client(t *testing.T) {
	got := CreateManager()
	rw := httptest.NewRecorder()
	form := url.Values{}
	form.Set("grant_type", "client_credentials")
	request := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	got.ProcessTokenEP(rw, request)
	println(rw.Code)
	toc := Tokens{}
	_ = json.NewDecoder(rw.Body).Decode(&toc)
	println(toc.RefreshToken)

	rw = httptest.NewRecorder()
	form = url.Values{}
	form.Set("grant_type", "refresh_token")
	form.Set("refresh_token", toc.AccessToken)
	request = httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set("authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")
	got.ProcessTokenEP(rw, request)
	println(rw.Code)
	toc = Tokens{}
	_ = json.NewDecoder(rw.Body).Decode(&toc)
	println(toc.RefreshToken)
}

func CreateManager() sdk.IManager {
	config := sdk.NewConfig("http://localhost:8080")
	config.RefreshTokenEntropy = 0
	private, public := util.GenerateRSAKeyPair()
	strategy := strategies.NewDefaultStrategy(private, public)
	sequence := CreateDefaultSequence()
	sequence = append(sequence, memdbstore.NewInMemoryDB(true))
	got := DefaultManager(config, strategy, sequence...)
	return got
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func TestMac(t *testing.T) {
	key := []byte("key")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte("some message"))
	expectedMAC := mac.Sum(nil)
	hexS := hex.EncodeToString(expectedMAC)
	println(hexS)
	toString := base64.URLEncoding.EncodeToString(expectedMAC)
	println(toString)
}
