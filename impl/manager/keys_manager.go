package manager

import (
	"encoding/json"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"gopkg.in/square/go-jose.v2"
	"log"
	"net/http"
)

func (d *DefaultManager) ProcessKeysEP(writer http.ResponseWriter, req *http.Request) {
	secrets, err := d.SecretStore.GetAllSecrets(req.Context())
	if err != nil {
		hint := sdkerror.ErrMisconfiguration.WithHint(err.Error())
		err := d.ErrorWriter.WriteJsonError(hint, nil, writer, req)
		if err != nil {
			d.ErrorStrategy(err, writer)
		}
		return
	}
	writer.Header().Add(sdk.HeaderContentType, sdk.ContentTypeJson)
	writer.WriteHeader(http.StatusOK)
	var publicKeys = jose.JSONWebKeySet{}
	for _, key := range secrets.Keys {
		publicKeys.Keys = append(publicKeys.Keys, key.Public())
	}
	err = json.NewEncoder(writer).Encode(publicKeys)
	if err != nil {
		log.Println(err)
	}
}
