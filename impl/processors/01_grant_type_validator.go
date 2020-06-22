package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultGrantTypeValidator struct {
}

func (d *DefaultGrantTypeValidator) HandleTokenEP(_ context.Context, request sdk.ITokenRequest, response sdk.ITokenResponse) sdk.IError {
	grantType := request.GetGrantType()
	client := response.GetClient()
	if !client.GetApprovedGrantTypes().Has(grantType) {
		return sdkerror.InvalidGrant.WithDescription("grant not approved")
	}
	return nil
}
