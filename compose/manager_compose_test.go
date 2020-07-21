package compose

import (
	"encoding/base64"
	"encoding/json"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/example/demosession"
	"github.com/identityOrg/oidcsdk/example/memdbstore"
	"github.com/identityOrg/oidcsdk/example/secretkey"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"golang.org/x/oauth2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestDefaultManager_Password(t *testing.T) {
	got := CreateManager()
	rw := httptest.NewRecorder()
	form := url.Values{}
	form.Set("grant_type", sdk.GrantResourceOwnerPassword)
	form.Set("username", "user")
	form.Set("password", "user")
	request := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set(sdk.HeaderAuthorization, "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set(sdk.HeaderContentType, sdk.ContentTypeUrlEncodedForm)
	got.ProcessTokenEP(rw, request)
	println(rw.Code)
}

func TestDefaultManager_Client(t *testing.T) {
	got := CreateManager()
	rw := httptest.NewRecorder()
	form := url.Values{}
	form.Set("grant_type", sdk.GrantClientCredentials)
	request := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set(sdk.HeaderAuthorization, "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set(sdk.HeaderContentType, sdk.ContentTypeUrlEncodedForm)
	got.ProcessTokenEP(rw, request)
	println(rw.Code)
	toc := Tokens{}
	_ = json.NewDecoder(rw.Body).Decode(&toc)
	println(toc.RefreshToken)

	rw = httptest.NewRecorder()
	form = url.Values{}
	form.Set("grant_type", sdk.GrantRefreshToken)
	form.Set("refresh_token", toc.AccessToken)
	request = httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	request.Header.Set(sdk.HeaderAuthorization, "Basic "+base64.StdEncoding.EncodeToString([]byte("client:client")))
	request.Header.Set(sdk.HeaderContentType, sdk.ContentTypeUrlEncodedForm)
	got.ProcessTokenEP(rw, request)
	println(rw.Code)
	toc = Tokens{}
	_ = json.NewDecoder(rw.Body).Decode(&toc)
	println(toc.RefreshToken)
}

func CreateManager() sdk.IManager {
	config := sdk.NewConfig("http://localhost:8080")
	config.RefreshTokenEntropy = 0
	strategy := strategies.NewDefaultStrategy()
	sequence := CreateDefaultSequence()
	demoStore := memdbstore.NewInMemoryDB(true)
	demoSessionManager := demosession.NewManager("some-secure-key", "demo-session")
	secretKeyStore := secretkey.NewDefaultMemorySecretStore()
	sequence = append(sequence, demoStore, demoSessionManager, secretKeyStore, strategy)
	manager := DefaultManager(config, sequence...)

	return manager
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var conf = &oauth2.Config{
	ClientID:     "client",
	ClientSecret: "client",
	Scopes:       []string{sdk.ScopeOpenid, "offline", "offline_access"},
	Endpoint: oauth2.Endpoint{
		AuthURL:   "http://localhost:8080/oauth2/auth",
		TokenURL:  "http://localhost:8080/oauth2/token",
		AuthStyle: oauth2.AuthStyleInHeader,
	},
	RedirectURL: "http://localhost:8080/callback",
}

func TestAuthorization(t *testing.T) {
	got := CreateManager()

	rw := httptest.NewRecorder()
	authCodeURL := conf.AuthCodeURL("ehfbwejwjewkjevkwevj")
	request := httptest.NewRequest(http.MethodGet, authCodeURL, nil)
	got.ProcessAuthorizationEP(rw, request)
	println(rw.Code)

}
