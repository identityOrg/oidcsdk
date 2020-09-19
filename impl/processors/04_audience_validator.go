package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultAudienceValidationProcessor struct {
}

func NewDefaultAudienceValidationProcessor() *DefaultAudienceValidationProcessor {
	return &DefaultAudienceValidationProcessor{}
}

func (d *DefaultAudienceValidationProcessor) HandleTokenEP(ctx context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == sdk.GrantResourceOwnerPassword || requestContext.GetGrantType() == sdk.GrantClientCredentials {
		requestedAudience := requestContext.GetRequestedAudience()
		client := requestContext.GetClient()
		if len(requestedAudience) == 0 {
			return nil
		} else {
			if client.GetApprovedScopes().Has(requestedAudience...) {
				requestContext.GetProfile().SetAudience(requestedAudience)
			} else {
				return sdkerror.ErrInvalidRequest.WithHint("un-approved audience requested")
			}
		}
	}
	return nil
}

func (d *DefaultAudienceValidationProcessor) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	requestedAudience := requestContext.GetRequestedAudience()
	client := requestContext.GetClient()
	if len(requestedAudience) == 0 {
		return nil
	} else {
		if client.GetApprovedScopes().Has(requestedAudience...) {
			requestContext.GetProfile().SetAudience(requestedAudience)
			return nil
		} else {
			return sdkerror.ErrInvalidRequest.WithHint("un-approved audience requested")
		}
	}
}
