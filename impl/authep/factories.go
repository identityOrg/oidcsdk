package authep

import (
	"github.com/google/uuid"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
	"oauth2-oidc-sdk/util"
	"strings"
	"time"
)

func DefaultAuthenticationRequestContextFactory(r *http.Request) (sdk.IAuthenticationRequestContext, sdk.IError) {
	if r.Method != http.MethodGet {
		return nil, sdkerror.ErrUnknownRequest.WithDescription("only HTTP method 'get' is supported")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrUnknownRequest.WithDescription(err.Error())
	}
	reqStruct := DefaultAuthenticationRequestContext{}
	form := r.Form

	reqStruct.RequestedScopes = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "scope"), " "))
	reqStruct.RequestedAudience = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "audience"), " "))
	reqStruct.ResponseType = util.RemoveEmpty(strings.Split(util.GetAndRemove(form, "response_type"), " "))
	reqStruct.RedirectURI = util.GetAndRemove(form, "redirect_uri")
	reqStruct.State = util.GetAndRemove(form, "state")
	reqStruct.ClientId = util.GetAndRemove(form, "client_id")
	reqStruct.Nonce = util.GetAndRemove(form, "nonce")

	reqStruct.Form = &form
	reqStruct.RequestID = uuid.New().String()
	reqStruct.RequestedAt = time.Now()

	reqStruct.Profile = sdk.NewRequestProfile()
	reqStruct.Claims = make(map[string]interface{})

	return &reqStruct, nil
}
