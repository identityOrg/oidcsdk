package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultResponseTypeValidator struct {
}

func NewDefaultResponseTypeValidator() *DefaultResponseTypeValidator {
	return &DefaultResponseTypeValidator{}
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
				return sdkerror.ErrUnsupportedResponseType.WithDescription("'authorization_code' grant not approved")
			}
			requestContext.GetProfile().SetGrantType(sdk.GrantAuthorizationCode)
		} else if responseType == sdk.ResponseTypeToken || responseType == sdk.ResponseTypeIdToken {
			if !approvedGrantTypes.Has(sdk.GrantImplicit) {
				return sdkerror.ErrUnsupportedResponseType.WithDebug("'implicit' grant not approved for client")
			}
			requestContext.GetProfile().SetGrantType(sdk.GrantImplicit)
			if responseType == sdk.ResponseTypeIdToken {
				nonce := requestContext.GetNonce()
				if nonce == "" {
					return sdkerror.ErrUnsupportedResponseType.WithHint("'nonce' is required for 'id_token' response")
				} else {
					requestContext.GetProfile().SetNonce(nonce)
				}
			}
		} else {
			return sdkerror.ErrUnsupportedResponseType.WithDebugf("un-known response type %s", responseType)
		}
	}
	return nil
}
