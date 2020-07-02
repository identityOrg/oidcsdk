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
		if responseType == sdk.ResponseTypeCode {
			if !approvedGrantTypes.Has(sdk.GrantAuthorizationCode) {
				return sdkerror.ErrInvalidGrant.WithDescription("'authorization_code' grant not approved")
			}
		} else if responseType == sdk.ResponseTypeToken || responseType == sdk.ResponseTypeIdToken {
			if !approvedGrantTypes.Has(sdk.GrantImplicit) {
				return sdkerror.ErrInvalidGrant.WithDebug("'implicit' grant not approved for client")
			}
		} else {
			return sdkerror.ErrUnsupportedResponseType.WithDebugf("un-known response type %s", responseType)
		}
	}
	return nil
}
