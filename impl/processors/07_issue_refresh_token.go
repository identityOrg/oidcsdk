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

func (d *DefaultRefreshTokenIssuer) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	refreshScope := requestContext.GetProfile().GetScope().Has(sdk.ScopeOfflineAccess)
	refreshGrant := requestContext.GetClient().GetApprovedGrantTypes().Has(sdk.GrantRefreshToken)
	if refreshScope || refreshGrant {
		token, signature := d.RefreshTokenStrategy.GenerateRefreshToken()
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		requestContext.IssueRefreshToken(token, signature, expiry)
	}
	return nil
}

func (d *DefaultRefreshTokenIssuer) Configure(config *sdk.Config, args ...interface{}) {
	for _, arg := range args {
		if us, ok := arg.(sdk.IRefreshTokenStrategy); ok {
			d.RefreshTokenStrategy = us
			break
		}
	}
	if d.RefreshTokenStrategy == nil {
		panic("failed to initialize DefaultRefreshTokenIssuer")
	}
	d.Lifespan = config.RefreshTokenLifespan
}
