package main

import (
	"net/http"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/compose"
	"oauth2-oidc-sdk/example/demosession"
	"oauth2-oidc-sdk/example/memdbstore"
	"oauth2-oidc-sdk/impl/strategies"
	"oauth2-oidc-sdk/util"
)

func main() {
	config := sdk.NewConfig("http://localhost:8080")
	config.RefreshTokenEntropy = 0
	private, public := util.GenerateRSAKeyPair()
	strategy := strategies.NewDefaultStrategy(private, public)
	sequence := compose.CreateDefaultSequence()
	sequence = append(sequence, memdbstore.NewInMemoryDB(true), demosession.NewManager("some-secure-key", "demo-session"))
	got := compose.DefaultManager(config, strategy, sequence...)
	compose.SetLoginPageHandler(got, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "text/html")
		_, _ = writer.Write([]byte(LoginPage))
	})
	compose.SetConsentPageHandler(got, func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		writer.Header().Set("Content-Type", "text/html")
		_, _ = writer.Write([]byte(ConsentPage))
	})

	http.HandleFunc("/token", got.ProcessTokenEP)
	http.HandleFunc("/authorize", got.ProcessAuthorizationEP)

	_ = http.ListenAndServe("127.0.0.1:8080", nil)
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
<form action="" method="post">
    <label for="u-name">Username</label>
    <input type="text" name="username" id="u-name">
    <label for="u-pass">Password</label>
    <input type="text" name="password" id="u-pass">
    <input type="submit">
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
