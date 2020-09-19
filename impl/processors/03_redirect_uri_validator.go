package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultRedirectURIValidator struct {
}

func NewDefaultRedirectURIValidator() *DefaultRedirectURIValidator {
	return &DefaultRedirectURIValidator{}
}

func (d *DefaultRedirectURIValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == sdk.GrantAuthorizationCode {
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

func (d *DefaultRedirectURIValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if client := requestContext.GetClient(); client == nil {
		return sdkerror.ErrUnauthorizedClient.WithDescription("client not resolved")
	} else {
		oidc := requestContext.GetRequestedScopes().Has(sdk.ScopeOpenid)
		redirectURI := requestContext.GetRedirectURI()
		if redirectURI == "" && !oidc {
			if len(client.GetRedirectURIs()) > 0 {
				derivedRedirectUri := client.GetRedirectURIs()[0]
				requestContext.SetRedirectURI(derivedRedirectUri)
				requestContext.GetProfile().SetRedirectURI(derivedRedirectUri)
				return nil
			}
		}
		for _, uri := range client.GetRedirectURIs() {
			if redirectURI == uri {
				requestContext.GetProfile().SetRedirectURI(redirectURI)
				return nil
			}
		}
		return sdkerror.ErrInvalidRequest.WithDescription("invalid redirect uri")
	}
}
