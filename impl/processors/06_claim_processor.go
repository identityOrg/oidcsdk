package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultClaimProcessor struct {
	UserStore sdk.IUserStore
}

func (d *DefaultClaimProcessor) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	openidReq := requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid)
	if openidReq && !requestContext.GetProfile().IsClient() {
		claims, err := d.UserStore.GetClaims(ctx, requestContext.GetProfile().GetUsername(), requestContext.GetProfile().GetScope(), nil)
		if err != nil {
			return sdkerror.ErrServerError.WithHint(err.Error())
		}
		contextClaim := requestContext.GetClaims()
		for s, i := range claims {
			contextClaim[s] = i
		}
	}
	return nil
}

func (d *DefaultClaimProcessor) HandleAuthEP(ctx context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if requestContext.GetProfile().GetScope().Has(sdk.ScopeOpenid) {
		if requestContext.GetResponseType().Has(sdk.ResponseTypeIdToken) {
			claims, err := d.UserStore.GetClaims(ctx, requestContext.GetProfile().GetUsername(), requestContext.GetProfile().GetScope(), nil)
			if err != nil {
				return sdkerror.ErrServerError.WithHint(err.Error())
			}
			contextClaim := requestContext.GetClaims()
			for s, i := range claims {
				contextClaim[s] = i
			}
		}
	}
	return nil
}

func (d *DefaultClaimProcessor) Configure(_ *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IUserStore); ok {
			d.UserStore = us
			break
		}
	}
	if d.UserStore == nil {
		panic("failed to initialize DefaultClaimProcessor")
	}
}

//var (
//	phoneScopeClaims   = strings.Split("phone_number_verified", ",")
//	emailScopeClaims   = strings.Split("phone_number,email,email_verified", ",")
//	profileScopeClaims = strings.Split("name,family_name,given_name,middle_name,nickname,preferred_username," +
//		"profile,picture,website,gender,birthdate,zoneinfo,locale,updated_at", ",")
//)
