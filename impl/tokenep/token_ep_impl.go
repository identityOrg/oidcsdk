package tokenep

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
	"oauth2-oidc-sdk/util"
	"strconv"
	"strings"
	"time"
)

type (
	DefaultTokenRequest struct {
		RequestID         string
		RequestedAt       time.Time
		State             string
		RedirectURI       string
		GrantType         string
		ClientId          string
		ClientSecret      string
		Username          string
		Password          string
		AuthorizationCode string
		RefreshToken      string
		RequestedScopes   sdk.Arguments
		RequestedAudience sdk.Arguments
		Form              *url.Values
	}
	DefaultTokenResponse struct {
		RequestID       string
		GrantedScopes   sdk.Arguments
		GrantedAudience sdk.Arguments
		Success         bool
		Client          sdk.IClient
		Profile         sdk.IProfile
		IssuedTokens    sdk.ITokens
		Error           sdk.IError
		Form            *url.Values
	}
)

func (d *DefaultTokenRequest) GetUsername() string {
	return d.Username
}

func (d *DefaultTokenRequest) GetPassword() string {
	return d.Password
}

func (d *DefaultTokenResponse) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultTokenResponse) GetGrantedScopes() sdk.Arguments {
	return d.GrantedScopes
}

func (d *DefaultTokenResponse) GetGrantedAudience() sdk.Arguments {
	return d.GrantedAudience
}

func (d *DefaultTokenResponse) GrantScope(scope string) {
	d.GrantedScopes = util.AppendUnique(d.GrantedScopes, scope)
}

func (d *DefaultTokenResponse) GrantAudience(audience string) {
	d.GrantedAudience = util.AppendUnique(d.GrantedAudience, audience)
}

func (d *DefaultTokenResponse) IsSuccess() bool {
	return d.Success
}

func (d *DefaultTokenResponse) SetSuccess(success bool) {
	d.Success = success
}

func (d *DefaultTokenResponse) GetClient() sdk.IClient {
	return d.Client
}

func (d *DefaultTokenResponse) SetClient(client sdk.IClient) {
	d.Client = client
}

func (d *DefaultTokenResponse) GetProfile() sdk.IProfile {
	return d.Profile
}

func (d *DefaultTokenResponse) SetProfile(profile sdk.IProfile) {
	d.Profile = profile
}

func (d *DefaultTokenResponse) GetIssuedTokens() sdk.ITokens {
	return d.IssuedTokens
}

func (d *DefaultTokenResponse) IssueTokens(tokens sdk.ITokens) {
	d.IssuedTokens = tokens
}

func (d *DefaultTokenResponse) GetError() sdk.IError {
	return d.Error
}

func (d *DefaultTokenResponse) SetError(err sdk.IError) {
	d.Error = err
}

func (d *DefaultTokenResponse) GetForm() *url.Values {
	return d.Form
}

// starting request

func (d *DefaultTokenRequest) GetRequestID() string {
	return d.RequestID
}

func (d *DefaultTokenRequest) GetRequestedAt() time.Time {
	return d.RequestedAt
}

func (d *DefaultTokenRequest) GetState() string {
	return d.State
}

func (d *DefaultTokenRequest) GetRedirectURI() string {
	return d.RedirectURI
}

func (d *DefaultTokenRequest) GetGrantType() string {
	return d.GrantType
}

func (d *DefaultTokenRequest) GetClientId() string {
	return d.ClientId
}

func (d *DefaultTokenRequest) GetClientSecret() string {
	return d.ClientSecret
}

func (d *DefaultTokenRequest) GetAuthorizationCode() string {
	return d.AuthorizationCode
}

func (d *DefaultTokenRequest) GetRefreshToken() string {
	return d.RefreshToken
}

func (d *DefaultTokenRequest) GetRequestedScopes() sdk.Arguments {
	return d.RequestedScopes
}

func (d *DefaultTokenRequest) GetRequestedAudience() sdk.Arguments {
	return d.RequestedAudience
}

func (d *DefaultTokenRequest) GetForm() *url.Values {
	return d.Form
}

func DefaultTokenRequestFactory(r *http.Request) (sdk.ITokenRequest, sdk.IError) {
	if r.Method != http.MethodPost {
		return nil, sdkerror.InvalidRequest.WithDescription("only HTTP method post supported")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, sdkerror.InvalidRequest.WithDescription(err.Error())
	}
	reqStruct := DefaultTokenRequest{}
	form := r.PostForm

	reqStruct.RequestedScopes = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "scope"), " "))
	reqStruct.RequestedAudience = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "audience"), " "))
	reqStruct.RefreshToken = util.GetAndRemove(form, "refresh_token")
	reqStruct.AuthorizationCode = util.GetAndRemove(form, "authorization_code")
	reqStruct.GrantType = util.GetAndRemove(form, "grant_type")
	reqStruct.RedirectURI = util.GetAndRemove(form, "redirect_uri")
	reqStruct.Username = util.GetAndRemove(form, "username")
	reqStruct.Password = util.GetAndRemove(form, "password")
	reqStruct.State = util.GetAndRemove(form, "state")
	reqStruct.ClientId = util.GetAndRemove(form, "client_id")
	reqStruct.ClientSecret = util.GetAndRemove(form, "client_secret")

	// check basic authorization
	authorization := r.Header.Get("authorization")
	parts := strings.SplitN(authorization, " ", 2)
	if strings.ToLower(parts[0]) == "basic" && len(parts) == 2 {
		decodeString, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			return nil, sdkerror.InvalidRequest.WithDescription(err.Error())
		}
		parts = strings.SplitN(string(decodeString), ":", 2)
		if len(parts) != 2 {
			return nil, sdkerror.InvalidRequest.WithDescription("invalid basic authorization header")
		}
		reqStruct.ClientId = parts[0]
		reqStruct.ClientSecret = parts[1]
	}

	reqStruct.RequestID = uuid.New().String()
	reqStruct.RequestedAt = time.Now()
	return &reqStruct, nil
}

func DefaultTokenResponseFactory(request sdk.ITokenRequest) (sdk.ITokenResponse, sdk.IError) {
	return &DefaultTokenResponse{
		RequestID: request.GetRequestID(),
	}, nil
}

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
		values["expiry"] = strconv.Itoa(int(tokens.GetAccessTokenExpiry().Seconds()))
		values["type"] = "bearer"
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