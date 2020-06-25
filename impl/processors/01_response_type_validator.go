package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultResponseTypeValidator struct {
}

func (d *DefaultResponseTypeValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) (sdk.IError, sdk.Result) {
	responseTypes := requestContext.GetResponseType()
	if len(responseTypes) == 0 {
		return sdkerror.InvalidRequest.WithDescription("no response type provided"), sdk.ResultNoOperation
	}
	approvedGrantTypes := requestContext.GetClient().GetApprovedGrantTypes()

	for _, responseType := range responseTypes {
		if responseType == "code" {
			if !approvedGrantTypes.Has("authorization_code") {
				return sdkerror.InvalidGrant.WithDescription("'authorization_code' grant not approved"), sdk.ResultNoOperation
			}
		} else if responseType == "token" || responseType == "id_token" {
			if !approvedGrantTypes.Has("implicit") {
				return sdkerror.InvalidGrant.WithDescription("'implicit' grant not approved"), sdk.ResultNoOperation
			}
		} else {
			return sdkerror.InvalidGrant.WithDescription("un-known response type"), sdk.ResultNoOperation
		}
	}
	return nil, sdk.ResultNoOperation
}
