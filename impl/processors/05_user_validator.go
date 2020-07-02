package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/impl/sdkerror"
	"oauth2-oidc-sdk/util"
	"strings"
)

type DefaultUserValidator struct {
	UserStore             sdk.IUserStore
	ClientStore           sdk.IClientStore
	GlobalConsentRequired bool
}

func (d *DefaultUserValidator) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	session := requestContext.GetUserSession()
	isOidc := requestContext.GetRequestedScopes().Has(sdk.ScopeOpenid)
	if !isOidc {
		if session.GetUsername() == "" {
			return sdkerror.ErrLoginRequired
		}
		if d.GlobalConsentRequired {
			if d.UserStore.IsConsentRequired(ctx, session.GetUsername(), requestContext.GetClientID(), requestContext.GetRequestedScopes()) {
				if !session.IsConsentSubmitted() {
					return sdkerror.ErrConsentRequired
				} else {
					for _, s := range session.GetApprovedScopes() {
						requestContext.GrantScope(s)
					}
				}
			}
		}
		return nil
	} else {
		var prompt sdk.Arguments = util.RemoveEmpty(strings.Split(util.GetAndRemove(*requestContext.GetForm(), "prompt"), " "))
		if prompt.Has("none") {
			if len(prompt) > 1 {
				return sdkerror.ErrInvalidRequest.WithHint("'prompt=none' can not be combined")
			}
		}
		if prompt.Has("login") && !session.IsLoginDone() {
			return sdkerror.ErrLoginRequired
		}
		if prompt.Has("consent") && !session.IsConsentSubmitted() {
			return sdkerror.ErrConsentRequired
		}
	}

	return nil
}

func (d *DefaultUserValidator) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	grantType := requestContext.GetGrantType()
	if grantType == sdk.GrantResourceOwnerPassword {
		username := requestContext.GetUsername()
		password := requestContext.GetPassword()
		err := d.UserStore.Authenticate(ctx, username, []byte(password))
		if err != nil {
			return sdkerror.ErrInvalidGrant.WithDescription("user authentication failed")
		}
		profile := d.UserStore.FetchUserProfile(ctx, username)
		profile.SetScope(requestContext.GetGrantedScopes())
		profile.SetAudience(requestContext.GetGrantedAudience())
		requestContext.SetProfile(profile)
	} else if grantType == sdk.GrantClientCredentials {
		profile := d.ClientStore.FetchClientProfile(ctx, requestContext.GetClientID())
		profile.SetScope(requestContext.GetGrantedScopes())
		profile.SetAudience(requestContext.GetGrantedAudience())
		requestContext.SetProfile(profile)
	}
	return nil
}

func (d *DefaultUserValidator) Configure(_ interface{}, config *sdk.Config, args ...interface{}) {
	d.GlobalConsentRequired = config.GlobalConsentRequired
	for _, arg := range args {
		if us, ok := arg.(sdk.IUserStore); ok {
			d.UserStore = us
		}
		if cs, ok := arg.(sdk.IClientStore); ok {
			d.ClientStore = cs
		}
	}
	if d.UserStore == nil || d.ClientStore == nil {
		panic("failed to init DefaultUserValidator")
	}
}
