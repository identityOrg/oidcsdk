package tokenep

import (
	"encoding/json"
	"errors"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"strconv"
)

func DefaultTokenResponseWriter(response sdk.ITokenResponse, w http.ResponseWriter) error {
	w.WriteHeader(200)
	w.Header().Set("content-type", "application/json")
	tokens := response.GetIssuedTokens()
	if tokens == nil {
		return errors.New("can not render nil tokens from token endpoint")
	}
	values := make(map[string]string)
	if tokens.GetAccessToken() != "" {
		values["access_token"] = tokens.GetAccessToken()
		values["expires_in"] = strconv.Itoa(int(tokens.GetAccessTokenExpiry().Seconds()))
		values["token_type"] = "bearer"
		values["scope"] = response.GetGrantedScopes().String()
	}
	if tokens.GetRefreshToken() != "" {
		values["refresh_token"] = tokens.GetRefreshToken()
	}
	if tokens.GetIDToken() != "" {
		values["id_token"] = tokens.GetIDToken()
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
