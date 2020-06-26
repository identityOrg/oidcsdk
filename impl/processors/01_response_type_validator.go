package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultResponseTypeValidator struct {
}

func (d *DefaultResponseTypeValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	responseTypes := requestContext.GetResponseType()
	if len(responseTypes) == 0 {
		return sdkerror.ErrInvalidRequest.WithDescription("no response type provided")
	}
	approvedGrantTypes := requestContext.GetClient().GetApprovedGrantTypes()

	for _, responseType := range responseTypes {
		if responseType == "code" {
			if !approvedGrantTypes.Has("authorization_code") {
				return sdkerror.ErrInvalidGrant.WithDescription("'authorization_code' grant not approved")
			}
		} else if responseType == "token" || responseType == "id_token" {
			if !approvedGrantTypes.Has("implicit") {
				return sdkerror.ErrInvalidGrant.WithDebug("'implicit' grant not approved for client")
			}
		} else {
			return sdkerror.ErrUnsupportedResponseType.WithDebugf("un-known response type %s", responseType)
		}
	}
	return nil
}
