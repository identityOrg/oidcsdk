package tokenep

import (
	"encoding/json"
	"errors"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"strconv"
	"time"
)

func DefaultTokenResponseWriter(response sdk.ITokenRequestContext, w http.ResponseWriter) error {
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	tokens := response.GetIssuedTokens()
	if tokens == nil {
		return errors.New("can not render nil tokens from token endpoint")
	}
	values := make(map[string]string)
	if tokens.AccessToken != "" {
		values["access_token"] = tokens.AccessToken
		values["expires_in"] = strconv.Itoa(tokens.AccessTokenExpiry.UTC().Round(time.Second).Second())
		values["token_type"] = "bearer"
		values["scope"] = response.GetGrantedScopes().String()
	}
	if tokens.RefreshToken != "" {
		values["refresh_token"] = tokens.RefreshToken
	}
	if tokens.IDToken != "" {
		values["id_token"] = tokens.IDToken
	}
	err := json.NewEncoder(w).Encode(values)
	return err
}

func DefaultTokenErrorWriter(pError sdk.IError, w http.ResponseWriter) error {
	w.WriteHeader(pError.GetStatusCode())
	w.Header().Set("content-type", "application/json")
	values := make(map[string]string)
	values["error"] = pError.GetErrorCode()
	values["error_description"] = pError.GetDescription()
	values["error_uri"] = pError.GetErrorURL()

	err := json.NewEncoder(w).Encode(values)
	return err
}
