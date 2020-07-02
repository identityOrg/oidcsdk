package main

import (
	"html/template"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/compose"
	"oauth2-oidc-sdk/example/demosession"
	"oauth2-oidc-sdk/example/memdbstore"
	"oauth2-oidc-sdk/impl/strategies"
	"oauth2-oidc-sdk/util"
	"time"
)

func main() {
	config := sdk.NewConfig("http://localhost:8080")
	config.RefreshTokenEntropy = 0
	private, public := util.GenerateRSAKeyPair()
	strategy := strategies.NewDefaultStrategy(private, public)
	sequence := compose.CreateDefaultSequence()
	demoStore := memdbstore.NewInMemoryDB(true)
	demoSessionManager := demosession.NewManager("some-secure-key", "demo-session")
	sequence = append(sequence, demoStore, demoSessionManager)
	got := compose.DefaultManager(config, strategy, sequence...)
	compose.SetLoginPageHandler(got, renderLogin)
	compose.SetConsentPageHandler(got, renderConsent)

	http.HandleFunc("/token", got.ProcessTokenEP)
	http.HandleFunc("/authorize", got.ProcessAuthorizationEP)
	http.HandleFunc("/login", processLogin(demoStore, demoSessionManager))

	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}

func processLogin(demoStore *memdbstore.InMemoryDB, manager *demosession.Manager) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
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
	}
}

func renderConsent(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	_, _ = writer.Write([]byte(ConsentPage))
}

func renderLogin(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
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
