package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultRedirectURIValidator struct {
}

func (d *DefaultRedirectURIValidator) HandleTokenEP(_ context.Context, request sdk.ITokenRequest, response sdk.ITokenResponse) sdk.IError {
	if request.GetGrantType() == "authorization_code" {
		if profile := response.GetProfile(); profile != nil {
			if profile.GetRedirectURI() != "" && request.GetRedirectURI() != profile.GetRedirectURI() {
				return sdkerror.InvalidGrant.WithDescription("redirect URI mismatch")
			} else {
				return nil
			}
		}
		return sdkerror.InvalidGrant.WithDescription("redirect URI validation failed")
	} else {
		return nil
	}
}

func (d *DefaultRedirectURIValidator) HandleAuthEP(_ context.Context, request sdk.IAuthenticationRequest, response sdk.IAuthenticationResponse) sdk.IError {
	if client := response.GetClient(); client == nil {
		return sdkerror.UnAuthorizedClient.WithDescription("client not resolved")
	} else {
		oidc := request.GetRequestedScopes().Has("openid")
		redirectURI := request.GetRedirectURI()
		if redirectURI == "" && !oidc {
			if len(client.GetRedirectURIs()) > 0 {
				request.SetRedirectURI(client.GetRedirectURIs()[0])
				return nil
			}
		}
		for _, uri := range client.GetRedirectURIs() {
			if redirectURI == uri {
				return nil
			}
		}
		return sdkerror.InvalidRequest.WithDescription("invalid redirect uri")
	}
}

//func (d *DefaultRedirectURIValidator) Configure(_ sdk.IManager, config sdk.Config, arg ...interface{}) {
//	panic("implement me")
//}
