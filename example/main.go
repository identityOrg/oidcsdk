package main

import (
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/compose"
	"github.com/identityOrg/oidcsdk/example/demosession"
	"github.com/identityOrg/oidcsdk/example/memdbstore"
	"github.com/identityOrg/oidcsdk/example/secretkey"
	"github.com/identityOrg/oidcsdk/impl/middleware"
	"github.com/identityOrg/oidcsdk/impl/strategies"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	config := sdk.NewConfig("http://localhost:8080")
	config.RefreshTokenEntropy = 0
	strategy := strategies.NewDefaultStrategy()
	sequence := compose.CreateDefaultSequence()
	demoStore := memdbstore.NewInMemoryDB(true)
	demoSessionManager := demosession.NewManager("some-secure-key", "demo-session")
	secretKeyStore := secretkey.NewDefaultMemorySecretStore()
	sequence = append(sequence, demoStore, demoSessionManager, secretKeyStore, strategy)
	manager := compose.DefaultManager(config, sequence...)
	compose.SetLoginPageHandler(manager, renderLogin)
	compose.SetConsentPageHandler(manager, renderConsent)

	router := compose.CreateNewRouter(manager)
	router.Methods(http.MethodPost).Path("/login").Handler(middleware.NoCache(processLogin(demoStore, demoSessionManager)))

	log.Println(http.ListenAndServe("localhost:8080", router))
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
