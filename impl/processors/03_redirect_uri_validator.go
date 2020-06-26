package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultRedirectURIValidator struct {
}

func (d *DefaultRedirectURIValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == "authorization_code" {
		if profile := requestContext.GetProfile(); profile != nil {
			if profile.GetRedirectURI() != "" && requestContext.GetRedirectURI() != profile.GetRedirectURI() {
				return sdkerror.ErrInvalidGrant.WithDescription("redirect URI mismatch")
			} else {
				return nil
			}
		}
		return sdkerror.ErrInvalidGrant.WithDescription("redirect URI validation failed")
	} else {
		return nil
	}
}

func (d *DefaultRedirectURIValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) (sdk.IError, sdk.Result) {
	if client := requestContext.GetClient(); client == nil {
		return sdkerror.ErrUnauthorizedClient.WithDescription("client not resolved"), sdk.ResultNoOperation
	} else {
		oidc := requestContext.GetRequestedScopes().Has("openid")
		redirectURI := requestContext.GetRedirectURI()
		if redirectURI == "" && !oidc {
			if len(client.GetRedirectURIs()) > 0 {
				requestContext.SetRedirectURI(client.GetRedirectURIs()[0])
				return nil, sdk.ResultNoOperation
			}
		}
		for _, uri := range client.GetRedirectURIs() {
			if redirectURI == uri {
				return nil, sdk.ResultNoOperation
			}
		}
		return sdkerror.ErrInvalidRequest.WithDescription("invalid redirect uri"), sdk.ResultNoOperation
	}
}

//func (d *DefaultRedirectURIValidator) Configure(_ sdk.IManager, config *sdk.Config, arg ...interface{}) {
//	panic("implement me")
//}
