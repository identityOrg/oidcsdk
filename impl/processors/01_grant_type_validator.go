package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultGrantTypeValidator struct {
}

func (d *DefaultGrantTypeValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	grantType := requestContext.GetGrantType()
	client := requestContext.GetClient()
	if !client.GetApprovedGrantTypes().Has(grantType) {
		return sdkerror.InvalidGrant.WithDescription("grant not approved")
	}
	if grantType == "client_credentials" && client.IsPublic() {
		return sdkerror.UnSupportedGrantType.WithDescription("'client_credentials' grant not allowed for public client")
	}
	if grantType == "password" && client.IsPublic() {
		return sdkerror.UnSupportedGrantType.WithDescription("'password' grant not allowed for public client")
	}
	return nil
}