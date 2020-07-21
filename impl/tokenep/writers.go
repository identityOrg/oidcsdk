package tokenep

import (
	"encoding/json"
	"fmt"
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
	"time"
)

func DefaultTokenResponseWriter(response sdk.ITokenRequestContext, w http.ResponseWriter, _ *http.Request) error {
	w.Header().Set(sdk.HeaderContentType, sdk.ContentTypeJson)
	tokens := response.GetIssuedTokens()
	values := make(map[string]string)
	if tokens.AccessToken != "" {
		values["access_token"] = tokens.AccessToken
		expiry := tokens.AccessTokenExpiry.Unix() - time.Now().Unix()
		values["expires_in"] = fmt.Sprintf("%d", expiry)
		values["token_type"] = "bearer"
		values["scope"] = response.GetProfile().GetScope().String()
	}
	if tokens.RefreshToken != "" {
		values["refresh_token"] = tokens.RefreshToken
	}
	if tokens.IDToken != "" {
		values["id_token"] = tokens.IDToken
	}
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(values)
	return err
}

func DefaultJsonErrorWriter(pError sdk.IError, w http.ResponseWriter, _ *http.Request) error {
	w.Header().Set(sdk.HeaderContentType, sdk.ContentTypeJson)
	w.WriteHeader(pError.GetStatusCode())

	err := json.NewEncoder(w).Encode(pError)
	return err
}
