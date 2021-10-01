package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultGrantTypeValidator struct {
}

func NewDefaultGrantTypeValidator() *DefaultGrantTypeValidator {
	return &DefaultGrantTypeValidator{}
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
	requestContext.GetProfile().SetGrantType(grantType)
	return nil
}
