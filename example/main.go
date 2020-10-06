package main

import (
	"github.com/gorilla/mux"
	sdk "github.com/identityOrg/oidcsdk"
	config2 "github.com/identityOrg/oidcsdk/example/config"
	"github.com/identityOrg/oidcsdk/example/demosession"
	"github.com/identityOrg/oidcsdk/example/memdbstore"
	"github.com/identityOrg/oidcsdk/impl/middleware"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	config := sdk.NewConfig("http://localhost:8081")
	config.RefreshTokenEntropy = 0
	demoConfig := &config2.DemoConfig{
		SessionEncKey:     "some-secure-key",
		SessionCookieName: "demo-session",
	}
	newManager := ComposeNewManager(config, true, demoConfig)
	newManager.SetErrorStrategy(strategies.DefaultLoggingErrorStrategy)
	newManager.SetLoginPageHandler(renderLogin)
	newManager.SetConsentPageHandler(renderConsent)

	router := CreateNewRouter(newManager)
	sessionManager := ComposeSessionStore(demoConfig)
	store := ComposeDemoStore(demoConfig, true)
	router.Methods(http.MethodPost).Path("/login").Handler(middleware.NoCache(processLogin(store, sessionManager)))

	log.Println(http.ListenAndServe("localhost:8081", router))
}

func CreateNewRouter(sdkManager sdk.IManager) *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/oauth2").Subrouter()
	subRouter.Use(middleware.NoCache)
	subRouter.Methods(http.MethodPost).Path("/token").HandlerFunc(sdkManager.ProcessTokenEP)
	subRouter.Methods(http.MethodGet).Path("/authorize").HandlerFunc(sdkManager.ProcessAuthorizationEP)
	subRouter.Methods(http.MethodPost).Path("/introspection").HandlerFunc(sdkManager.ProcessIntrospectionEP)
	subRouter.Methods(http.MethodPost).Path("/revocation").HandlerFunc(sdkManager.ProcessRevocationEP)
	subRouter.Methods(http.MethodGet).Path("/keys").HandlerFunc(sdkManager.ProcessKeysEP)
	subRouter.Methods(http.MethodGet).Path("/me").HandlerFunc(sdkManager.ProcessUserInfoEP)
	router.Methods(http.MethodGet).Path(sdk.UrlOidcDiscovery).Handler(middleware.NoCache(http.HandlerFunc(sdkManager.ProcessDiscoveryEP)))
	return router
}

func processLogin(demoStore *memdbstore.InMemoryDB, manager *demosession.Manager) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		username := request.PostForm.Get("username")
		password := request.PostForm.Get("password")
		err := demoStore.Authenticate(request.Context(), username, []byte(password))
		if err != nil {
			renderLogin(writer, request)
		} else {
			sess := demosession.DefaultSession{}
			now := time.Now()
			sess.LoginTime = &now
			sess.Username = username
			requestUrl := request.PostForm.Get("request")
			err = manager.StoreUserSession(writer, request, sess)
			if err != nil {
				writer.WriteHeader(500)
				_, _ = writer.Write([]byte(err.Error()))
			} else {
				http.Redirect(writer, request, requestUrl, http.StatusFound)
			}
		}
	})
}

func renderConsent(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_, _ = writer.Write([]byte(ConsentPage))
}

func renderLogin(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = template.Must(template.New("login").Parse(LoginPage)).Execute(writer, request.URL.String())
}

const (
	LoginPage = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Login</title>
</head>
<body>
<form action="/login" method="post">
    <label for="u-name">Username</label>
    <input type="text" name="username" id="u-name">
    <label for="u-pass">Password</label>
    <input type="text" name="password" id="u-pass">
    <input type="submit">
    <input type="hidden" name="request" value="{{.}}">
</form>
</body>
</html>
`
	ConsentPage = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Login</title>
</head>
<body>
<form action="" method="post">
    <p>Consent Required</p>
</form>
</body>
</html>
`
)
