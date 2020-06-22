package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultClientAuthenticationProcessor struct {
	ClientStore sdk.IClientStore
}

func (d *DefaultClientAuthenticationProcessor) Configure(_ interface{}, _ sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if cs, ok := arg.(sdk.IClientStore); ok {
			d.ClientStore = cs
			break
		}
	}
	if d.ClientStore == nil {
		panic("failed in init of DefaultClientAuthenticationProcessor")
	}
}

func (d *DefaultClientAuthenticationProcessor) HandleAuthEP(_ context.Context, request sdk.IAuthenticationRequest, response sdk.IAuthenticationResponse) sdk.IError {
	clientId := request.GetClientID()
	if clientId == "" {
		return sdkerror.InvalidClient.WithDescription("client id not found in request")
	}
	client, err := d.ClientStore.GetClient(clientId)
	if err != nil {
		return sdkerror.InvalidClient.WithDescription(err.Error())
	}
	response.SetClient(client)
	return nil
}

func (d *DefaultClientAuthenticationProcessor) HandleTokenEP(_ context.Context, request sdk.ITokenRequest, response sdk.ITokenResponse) sdk.IError {
	clientId := request.GetClientID()
	if clientId == "" {
		return sdkerror.InvalidClient.WithDescription("client id not found in request")
	}
	clientSecret := request.GetClientSecret()
	client, err := d.ClientStore.GetClient(clientId)
	if err != nil {
		return sdkerror.InvalidClient.WithDescription(err.Error())
	}
	if clientSecret == "" && client.IsPublic() {
		response.SetClient(client)
		return nil
	}

	//todo handle secret encryption
	if client.GetSecret() == clientSecret {
		response.SetClient(client)
		return nil
	}
	return sdkerror.InvalidClient.WithDescription("could not authenticated client")
}
