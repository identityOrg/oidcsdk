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
		return sdkerror.UnAuthorizedClient.WithDescription("client not resolved")
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
				return sdkerror.InvalidScope.WithDescription("un-approved or invalid scope requested")
			}
		}
		return nil
	}
}

func (d *DefaultScopeValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	grantType := requestContext.GetGrantType()
	if grantType == "authorization_code" {
		profile := requestContext.GetProfile()
		if profile.GetScope().MatchesExact(requestContext.GetRequestedScopes()...) {
			for _, s := range requestContext.GetRequestedScopes() {
				requestContext.GrantScope(s)
			}
			return nil
		} else {
			return sdkerror.InvalidScope.WithDescription("mismatch in requested scope")
		}
	} else if grantType == "password" || grantType == "client_credentials" {
		approvedScopes := requestContext.GetClient().GetApprovedScopes()
		for _, requestedScope := range requestContext.GetRequestedScopes() {
			if approvedScopes.Has(requestedScope) {
				requestContext.GrantScope(requestedScope)
			}
		}
	} else if grantType == "refresh_token" {
		scope := requestContext.GetProfile().GetScope()
		for _, s := range scope {
			requestContext.GrantScope(s)
		}
	}
	return nil
}
