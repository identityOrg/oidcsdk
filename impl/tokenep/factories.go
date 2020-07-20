package tokenep

import (
	"encoding/base64"
	"github.com/google/uuid"
	"net/http"
	sdk "oidcsdk"
	"oidcsdk/impl/sdkerror"
	"oidcsdk/util"
	"strings"
	"time"
)

func DefaultTokenRequestContextFactory(r *http.Request) (sdk.ITokenRequestContext, sdk.IError) {
	if r.Method != http.MethodPost {
		return nil, sdkerror.ErrInvalidRequest.WithDescription("only HTTP method post supported")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrInvalidRequest.WithDescription(err.Error())
	}
	reqStruct := DefaultTokenRequestContext{}
	form := r.PostForm

	reqStruct.RequestedScopes = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "scope"), " "))
	reqStruct.RequestedAudience = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "audience"), " "))
	reqStruct.RefreshToken = util.GetAndRemove(form, "refresh_token")
	reqStruct.AuthorizationCode = util.GetAndRemove(form, "code")
	reqStruct.GrantType = util.GetAndRemove(form, "grant_type")
	reqStruct.RedirectURI = util.GetAndRemove(form, "redirect_uri")
	reqStruct.Username = util.GetAndRemove(form, "username")
	reqStruct.Password = util.GetAndRemove(form, "password")
	reqStruct.State = util.GetAndRemove(form, "state")
	reqStruct.ClientId = util.GetAndRemove(form, "client_id")
	reqStruct.ClientSecret = util.GetAndRemove(form, "client_secret")

	// check basic authorization
	authorization := r.Header.Get(sdk.HeaderAuthorization)
	parts := strings.SplitN(authorization, " ", 2)
	if strings.ToLower(parts[0]) == "basic" && len(parts) == 2 {
		decodeString, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			return nil, sdkerror.ErrInvalidRequest.WithDescription(err.Error())
		}
		parts = strings.SplitN(string(decodeString), ":", 2)
		if len(parts) != 2 {
			return nil, sdkerror.ErrInvalidRequest.WithDescription("invalid basic authorization header")
		}
		reqStruct.ClientId = parts[0]
		reqStruct.ClientSecret = parts[1]
	}

	reqStruct.RequestID = uuid.New().String()
	reqStruct.RequestedAt = time.Now()

	reqStruct.Profile = sdk.NewRequestProfile()
	reqStruct.Claims = make(map[string]interface{})

	return &reqStruct, nil
}
