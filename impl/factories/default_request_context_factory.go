package factories

import (
	"github.com/google/uuid"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"github.com/identityOrg/oidcsdk/util"
	"net/http"
	"strings"
	"time"
)

type DefaultRequestContextFactory struct {
}

func (d DefaultRequestContextFactory) BuildUserInfoRequestContext(request *http.Request) (sdk.IUserInfoRequestContext, sdk.IError) {
	rContext := &DefaultUserInfoRequestContext{
		Claims: make(map[string]interface{}),
	}
	if request.Method != http.MethodGet {
		return nil, sdkerror.ErrInvalidRequest.WithHintf("http method %s is not supported", request.Method)
	}
	authHeader := request.Header.Get(sdk.HeaderAuthorization)
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, sdkerror.ErrInvalidRequest.WithHint("missing bearer token in 'Authorization' header")
	} else {
		rContext.BearerToken = authHeader[7:]
		return rContext, nil
	}
}

func NewDefaultRequestContextFactory() *DefaultRequestContextFactory {
	return &DefaultRequestContextFactory{}
}

func (d DefaultRequestContextFactory) BuildTokenRequestContext(request *http.Request) (sdk.ITokenRequestContext, sdk.IError) {
	if request.Method != http.MethodPost {
		return nil, sdkerror.ErrInvalidRequest.WithDescription("only HTTP method post supported")
	}
	err := request.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrInvalidRequest.WithDescription(err.Error())
	}
	reqStruct := DefaultTokenRequestContext{}
	form := request.PostForm

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
	if reqStruct.ClientId, reqStruct.ClientSecret, ok = request.BasicAuth(); !ok {
		return nil, sdkerror.ErrUnauthorizedClient.WithHint("client authorization basic header not found")
	}

	reqStruct.Form = &form

	reqStruct.RequestID = uuid.New().String()
	reqStruct.RequestedAt = time.Now()

	reqStruct.Profile = sdk.NewRequestProfile()
	reqStruct.Claims = make(map[string]interface{})

	return &reqStruct, nil
}

func (d DefaultRequestContextFactory) BuildAuthorizationRequestContext(request *http.Request) (sdk.IAuthenticationRequestContext, sdk.IError) {
	if request.Method != http.MethodGet {
		return nil, sdkerror.ErrUnknownRequest.WithDescription("only HTTP method 'get' is supported")
	}
	err := request.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrUnknownRequest.WithDescription(err.Error())
	}
	reqStruct := DefaultAuthenticationRequestContext{}
	form := request.Form

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

func (d DefaultRequestContextFactory) BuildRevocationRequestContext(request *http.Request) (sdk.IRevocationRequestContext, sdk.IError) {
	if request.Method != http.MethodPost {
		return nil, sdkerror.ErrInvalidRequest.WithHintf("http method %s is not allowed for revocation", request.Method)
	}
	err := request.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrInvalidRequest.WithHint(err.Error())
	}
	requestContext := DefaultRevocationRequestContext{}

	ok := false
	requestContext.ClientID, requestContext.ClientSecret, ok = request.BasicAuth()
	if !ok {
		return nil, sdkerror.ErrRequestUnauthorized.WithHint("missing basic authorization header")
	}

	requestContext.RequestID = uuid.New().String()
	requestContext.RequestedAt = time.Now()

	form := request.PostForm
	requestContext.Token = util.GetAndRemove(form, "token")
	requestContext.TokenTypeHint = util.GetAndRemove(form, "token_type_hint")
	requestContext.Form = &form

	return &requestContext, nil
}

func (d DefaultRequestContextFactory) BuildIntrospectionRequestContext(request *http.Request) (sdk.IIntrospectionRequestContext, sdk.IError) {
	if request.Method != http.MethodPost {
		return nil, sdkerror.ErrInvalidRequest.WithHintf("http method %s is not allowed for introspection", request.Method)
	}
	err := request.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrInvalidRequest.WithHint(err.Error())
	}
	requestContext := DefaultIntrospectionRequestContext{}

	ok := false
	requestContext.ClientID, requestContext.ClientSecret, ok = request.BasicAuth()
	if !ok {
		return nil, sdkerror.ErrRequestUnauthorized.WithHint("missing basic authorization header")
	}

	requestContext.RequestID = uuid.New().String()
	requestContext.RequestedAt = time.Now()

	form := request.PostForm
	requestContext.Token = util.GetAndRemove(form, "token")
	requestContext.TokenTypeHint = util.GetAndRemove(form, "token_type_hint")
	requestContext.Form = &form

	return &requestContext, nil
}
