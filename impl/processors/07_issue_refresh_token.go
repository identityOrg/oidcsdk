package processors

import (
	"context"
	sdk "github.com/identityOrg/oidcsdk"
	"time"
)

type DefaultRefreshTokenIssuer struct {
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
	Lifespan             time.Duration
}

func NewDefaultRefreshTokenIssuer(refreshTokenStrategy sdk.IRefreshTokenStrategy, config *sdk.Config) *DefaultRefreshTokenIssuer {
	return &DefaultRefreshTokenIssuer{RefreshTokenStrategy: refreshTokenStrategy, Lifespan: config.RefreshTokenLifespan}
}

func (d *DefaultRefreshTokenIssuer) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	refreshScope := requestContext.GetProfile().GetScope().Has(sdk.ScopeOfflineAccess)
	refreshGrant := requestContext.GetClient().GetApprovedGrantTypes().Has(sdk.GrantRefreshToken)
	if refreshScope || refreshGrant {
		token, signature := d.RefreshTokenStrategy.GenerateRefreshToken()
		expiry := requestContext.GetRequestedAt().Add(d.Lifespan).Round(time.Second)
		requestContext.IssueRefreshToken(token, signature, expiry)
	}
	return nil
}
