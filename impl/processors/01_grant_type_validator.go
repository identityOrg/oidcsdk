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
		return sdkerror.ErrInvalidGrant.WithDescription("grant not approved")
	}
	if grantType == sdk.GrantClientCredentials && client.IsPublic() {
		return sdkerror.ErrUnsupportedGrantType.WithDescription("'client_credentials' grant not allowed for public client")
	}
	if grantType == sdk.GrantResourceOwnerPassword && client.IsPublic() {
		return sdkerror.ErrUnsupportedGrantType.WithDescription("'password' grant not allowed for public client")
	}
	return nil
}
