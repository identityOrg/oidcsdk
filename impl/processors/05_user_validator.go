package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"github.com/identityOrg/oidcsdk/util"
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
	} else {
		var prompt sdk.Arguments = util.RemoveEmpty(strings.Split(util.GetAndRemove(*requestContext.GetForm(), "prompt"), " "))
		offlineAccess := requestContext.GetRequestedScopes().Has(sdk.ScopeOfflineAccess)
		promptNone := prompt.Has("none")
		promptLogin := prompt.Has("login")
		promptConsent := prompt.Has("consent")
		if promptNone {
			if offlineAccess {
				return sdkerror.ErrConsentRequired.WithHint("'prompt=none' can not be combined with 'offline_access'")
			}
			if promptLogin || promptConsent {
				return sdkerror.ErrInvalidRequest.WithHint("'prompt=none' can not be combined with other prompt")
			}
		}
		if promptLogin && !session.IsLoginDone() {
			return sdkerror.ErrLoginRequired
		}
		if promptConsent && !session.IsConsentSubmitted() && !d.GlobalConsentRequired {
			return sdkerror.ErrConsentRequired
		}

		if session.GetUsername() == "" {
			return sdkerror.ErrLoginRequired
		}
	}

	profile := requestContext.GetProfile()
	profile.SetUsername(session.GetUsername())

	if d.GlobalConsentRequired {
		if d.UserStore.IsConsentRequired(ctx, session.GetUsername(), requestContext.GetClientID(), requestContext.GetRequestedScopes()) {
			if !session.IsConsentSubmitted() {
				return sdkerror.ErrConsentRequired
			} else {
				profile.SetScope(session.GetApprovedScopes())
			}
		} else {
			profile.SetScope(requestContext.GetRequestedScopes())
		}
	} else {
		profile.SetScope(requestContext.GetRequestedScopes())
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
		profile := requestContext.GetProfile()
		user := d.UserStore.FetchUserProfile(ctx, username)
		profile.SetUsername(user.GetUsername())
	} else if grantType == sdk.GrantClientCredentials {
		profile := requestContext.GetProfile()
		client := d.ClientStore.FetchClientProfile(ctx, requestContext.GetClientID())
		profile.SetUsername(client.GetUsername())
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
