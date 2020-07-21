package manager

import (
	"encoding/json"
	sdk "github.com/identityOrg/oidcsdk"
	"gopkg.in/square/go-jose.v2"
	"log"
	"net/http"
)

func (d *DefaultManager) ProcessKeysEP(writer http.ResponseWriter, _ *http.Request) {
	secrets := d.SecretStore.GetAllSecrets()
	writer.Header().Add(sdk.HeaderContentType, sdk.ContentTypeJson)
	writer.WriteHeader(http.StatusOK)
	var publicKeys = jose.JSONWebKeySet{}
	for _, key := range secrets.Keys {
		publicKeys.Keys = append(publicKeys.Keys, key.Public())
	}
	log.Println(json.NewEncoder(writer).Encode(publicKeys))
}
