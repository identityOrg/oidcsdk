package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultRPILogoutIDTokenValidator struct {
	JWTValidator sdk.IJWTValidator
	ClientStore  sdk.IClientStore
}

func NewDefaultRPILogoutIDTokenValidator(JWTValidator sdk.IJWTValidator, clientStore sdk.IClientStore) *DefaultRPILogoutIDTokenValidator {
	return &DefaultRPILogoutIDTokenValidator{JWTValidator: JWTValidator, ClientStore: clientStore}
}

func (d *DefaultRPILogoutIDTokenValidator) HandleRPILogoutEP(ctx context.Context, requestContext sdk.IRPILogoutRequestContext) sdk.IError {
	if requestContext.GetIdTokenHint() != "" {
		clientId, username, err := d.JWTValidator.ValidateOwnJWTToken(ctx, requestContext.GetIdTokenHint())
		if err != nil {
			return sdkerror.ErrInvalidTokenFormat.WithHint(err.Error())
		}
		client, err := d.ClientStore.GetClient(ctx, clientId)
		if err != nil {
			return sdkerror.ErrInvalidClient.WithHint("client with id " + clientId + " not found: " + err.Error())
		}
		requestContext.SetClient(client)
		requestContext.SetUsername(username)
	}
	return nil
}
