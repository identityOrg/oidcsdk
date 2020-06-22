package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
)

type DefaultScopeValidator struct {
}

func (d *DefaultScopeValidator) HandleAuthEP(_ context.Context, request sdk.IAuthenticationRequest, response sdk.IAuthenticationResponse) sdk.IError {
	if client := response.GetClient(); client == nil {
		return sdkerror.UnAuthorizedClient.WithDescription("client not resolved")
	} else {
		requestedScopes := request.GetRequestedScopes()
		approvedScopes := client.GetApprovedScopes()
		for _, scope := range requestedScopes {
			found := false
			for _, approved := range approvedScopes {
				if scope == approved {
					found = true
				}
			}
			if !found {
				return sdkerror.InvalidScope.WithDescription("un-approved or invalid scope requested")
			}
		}
		return nil
	}
}

func (d *DefaultScopeValidator) HandleTokenEP(_ context.Context, request sdk.ITokenRequest, response sdk.ITokenResponse) sdk.IError {
	if request.GetGrantType() == "authorization_code" {
		profile := response.GetProfile()
		if profile.GetScope().MatchesExact(request.GetRequestedScopes()...) {
			for _, s := range request.GetRequestedScopes() {
				response.GrantScope(s)
			}
			return nil
		} else {
			return sdkerror.InvalidScope.WithDescription("mismatch in requested scope")
		}
	} else if request.GetGrantType() == "password" || request.GetGrantType() == "client_credentials" {
		approvedScopes := response.GetClient().GetApprovedScopes()
		for _, requestedScope := range request.GetRequestedScopes() {
			if approvedScopes.Has(requestedScope) {
				response.GrantScope(requestedScope)
			}
		}
	} else if request.GetGrantType() == "refresh_token" {
		scope := response.GetProfile().GetScope()
		for _, s := range scope {
			response.GrantScope(s)
		}
	}
	return nil
}
