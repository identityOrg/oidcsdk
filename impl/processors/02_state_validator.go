package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultStateValidator struct {
	StateParamMinimumEntropy int
}

func (d *DefaultStateValidator) Configure(config *sdk.Config, _ ...interface{}) {
	d.StateParamMinimumEntropy = config.StateParamMinimumEntropy
}

func (d *DefaultStateValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	if len(requestContext.GetState()) < d.StateParamMinimumEntropy {
		return sdkerror.ErrInsufficientEntropy.WithHintf("state parameter entropy is less then %d",
			d.StateParamMinimumEntropy)
	}
	requestContext.GetProfile().SetState(requestContext.GetState())
	return nil
}

func (d *DefaultStateValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == sdk.GrantAuthorizationCode {
		profile := requestContext.GetProfile()
		state := requestContext.GetState()
		if profile.GetState() != "" && profile.GetState() != state {
			return sdkerror.ErrInvalidState.WithHint("state doesnt match with previous value")
		}
	}
	return nil
}
