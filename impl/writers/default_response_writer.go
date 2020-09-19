package writers

import (
	"encoding/json"
	"errors"
	"fmt"
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type DefaultResponseWriter struct {
}

func NewDefaultResponseWriter() *DefaultResponseWriter {
	return &DefaultResponseWriter{}
}

func (*DefaultResponseWriter) WriteUserInfoResponse(requestContext sdk.IUserInfoRequestContext, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(sdk.HeaderContentType, sdk.ContentTypeJson)
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(requestContext.GetClaims())
	return err
}

func (*DefaultResponseWriter) WriteTokenResponse(requestContext sdk.ITokenRequestContext, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set(sdk.HeaderContentType, sdk.ContentTypeJson)
	tokens := requestContext.GetIssuedTokens()
	values := make(map[string]string)
	if tokens.AccessToken != "" {
		values["access_token"] = tokens.AccessToken
		expiry := tokens.AccessTokenExpiry.Unix() - time.Now().Unix()
		values["expires_in"] = fmt.Sprintf("%d", expiry)
		values["token_type"] = "bearer"
		values["scope"] = requestContext.GetProfile().GetScope().String()
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

func (*DefaultResponseWriter) WriteAuthorizationResponse(requestContext sdk.IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error {
	mode := requestContext.GetResponseMode()
	switch mode {
	case sdk.ResponseModeFragment:
		form := buildSuccessResponseForm(requestContext.GetIssuedTokens())
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.Fragment = form.Encode()
		http.Redirect(w, r, redirectUri.String(), http.StatusFound)
		return nil
	case sdk.ResponseModeQuery:
		form := buildSuccessResponseForm(requestContext.GetIssuedTokens())
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.RawQuery = form.Encode()
		http.Redirect(w, r, redirectUri.String(), http.StatusFound)
		return nil
	}
	return errors.New("invalid response mode")
}

func buildSuccessResponseForm(tokens sdk.Tokens) url.Values {
	form := url.Values{}
	if tokens.AccessToken != "" {
		form.Add("access_token", tokens.AccessToken)
		expiry := tokens.AccessTokenExpiry.Unix() - time.Now().Unix()
		form.Add("expires_in", fmt.Sprintf("%d", expiry))
		form.Add("type", "bearer")
	}
	if tokens.AuthorizationCode != "" {
		form.Add(sdk.ResponseTypeCode, tokens.AuthorizationCode)
	}
	if tokens.IDToken != "" {
		form.Add("id_token", tokens.IDToken)
	}
	return form
}

func (*DefaultResponseWriter) WriteIntrospectionResponse(requestContext sdk.IIntrospectionRequestContext, w http.ResponseWriter, r *http.Request) error {
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

	w.Header().Add(sdk.HeaderContentType, sdk.ContentTypeJson)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
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

func (*DefaultResponseWriter) WriteRevocationResponse(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	return nil
}
