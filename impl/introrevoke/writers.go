package introrevoke

import (
	"encoding/json"
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
	"strings"
)

func DefaultIntrospectionResponseWriter(requestContext sdk.IIntrospectionRequestContext, writer http.ResponseWriter, request *http.Request) error {
	var resp introspectionResponse
	if requestContext.IsActive() {
		profile := requestContext.GetProfile()
		resp = introspectionResponse{
			Active:    requestContext.IsActive(),
			Scope:     strings.Join(profile.GetScope(), " "),
			ClientID:  requestContext.GetClientID(),
			Username:  profile.GetUsername(),
			TokenType: requestContext.GetTokenType(),
			Expiry:    0,
			IssuedAT:  0,
			NotBefore: 0,
			Subject:   "",
			Audience:  profile.GetAudience(),
			Issuer:    "",
			ID:        "",
		}
	} else {
		resp = introspectionResponse{
			Active: requestContext.IsActive(),
		}
	}

	writer.Header().Add(sdk.HeaderContentType, sdk.ContentTypeJson)
	writer.WriteHeader(http.StatusOK)
	return json.NewEncoder(writer).Encode(resp)
}

type introspectionResponse struct {
	Active    bool     `json:"active"`
	Scope     string   `json:"scope,omitempty"`
	ClientID  string   `json:"client_id,omitempty"`
	Username  string   `json:"username,omitempty"`
	TokenType string   `json:"token_type,omitempty"`
	Expiry    int64    `json:"exp,omitempty"`
	IssuedAT  int64    `json:"iat,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
	Subject   string   `json:"sub,omitempty"`
	Audience  []string `json:"aud,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	ID        string   `json:"jti,omitempty"`
}

func DefaultRevocationResponseWriter(requestContext sdk.IRevocationRequestContext, writer http.ResponseWriter, request *http.Request) error {
	writer.WriteHeader(200)
	return nil
}
