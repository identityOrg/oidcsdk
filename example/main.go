package main

import (
	"net/http"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/compose"
	"oauth2-oidc-sdk/impl/demosession"
	"oauth2-oidc-sdk/impl/memdbstore"
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

	http.HandleFunc("/token", got.ProcessTokenEP)
	http.HandleFunc("/authorize", func(writer http.ResponseWriter, request *http.Request) {
		result := got.ProcessAuthorizationEP(writer, request)
		switch result {
		case sdk.ResultLoginRequired:
			writer.WriteHeader(200)
			_, _ = writer.Write([]byte("login required"))
			break
		case sdk.ResultConsentRequired:
			writer.WriteHeader(200)
			writer.Write([]byte("consent required"))
			break
		}
	})

	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}
