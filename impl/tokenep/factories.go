package tokenep

import (
	"github.com/google/uuid"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"github.com/identityOrg/oidcsdk/util"
	"net/http"
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
	var ok bool
	if reqStruct.ClientId, reqStruct.ClientSecret, ok = r.BasicAuth(); !ok {
		return nil, sdkerror.ErrUnauthorizedClient.WithHint("client authorization basic header not found")
	}

	reqStruct.Form = &form

	reqStruct.RequestID = uuid.New().String()
	reqStruct.RequestedAt = time.Now()

	reqStruct.Profile = sdk.NewRequestProfile()
	reqStruct.Claims = make(map[string]interface{})

	return &reqStruct, nil
}
