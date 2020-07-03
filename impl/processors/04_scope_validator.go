package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultScopeValidator struct {
}

func (d *DefaultScopeValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if client := requestContext.GetClient(); client == nil {
		return sdkerror.ErrUnauthorizedClient.WithDescription("client not resolved")
	} else {
		requestedScopes := requestContext.GetRequestedScopes()
		approvedScopes := client.GetApprovedScopes()
		for _, scope := range requestedScopes {
			found := false
			for _, approved := range approvedScopes {
				if scope == approved {
					found = true
				}
			}
			if !found {
				return sdkerror.ErrInvalidScope.WithDescription("un-approved or invalid scope requested")
			}
		}
		return nil
	}
}

func (d *DefaultScopeValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	grantType := requestContext.GetGrantType()
	requestedScopes := requestContext.GetRequestedScopes()
	if grantType == sdk.GrantAuthorizationCode {
		profile := requestContext.GetProfile()
		if profile.GetScope().MatchesExact(requestedScopes...) {
			return nil
		} else {
			return sdkerror.ErrInvalidScope.WithDescription("mismatch in requested scope")
		}
	} else if grantType == sdk.GrantResourceOwnerPassword || grantType == sdk.GrantClientCredentials {
		approvedScopes := requestContext.GetClient().GetApprovedScopes()
		if approvedScopes.Has(requestedScopes...) {
			requestContext.GetProfile().SetScope(requestedScopes)
		} else {
			return sdkerror.ErrInvalidScope.WithHint("at least one scope is un-approved")
		}
	}
	return nil
}
