package introrevoke

import (
	"github.com/google/uuid"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"github.com/identityOrg/oidcsdk/util"
	"net/http"
	"time"
)

func DefaultIntrospectionRequestContextFactory(request *http.Request) (sdk.IIntrospectionRequestContext, sdk.IError) {
	if request.Method != http.MethodPost {
		return nil, sdkerror.ErrInvalidRequest.WithHintf("http method %s is not allowed for introspection", request.Method)
	}
	err := request.ParseForm()
	if err != nil {
		return nil, sdkerror.ErrInvalidRequest.WithHint(err.Error())
	}
	requestContext := DefaultIntrospectionRequestContext{}

	ok := false
	requestContext.ClientID, requestContext.Secret, ok = request.BasicAuth()
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
