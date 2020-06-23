package processors

import (
	"context"
	sdk "oauth2-oidc-sdk"
	"time"
)

type DefaultRefreshTokenIssuer struct {
	RefreshTokenStrategy sdk.IRefreshTokenStrategy
	TokenStore           sdk.ITokenStore
	Lifespan             time.Duration
}

func (d *DefaultRefreshTokenIssuer) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	refreshScope := requestContext.GetGrantedScopes().Has("offline_access")
	refreshGrant := requestContext.GetClient().GetApprovedGrantTypes().Has("refresh_token")
	if refreshScope || refreshGrant {
		token, signature := d.RefreshTokenStrategy.GenerateRefreshToken()
		expiry := requestContext.GetRequestedAt().UTC().Add(d.Lifespan).Round(time.Second)
		requestContext.IssueRefreshToken(token, signature, expiry)
	}
	return nil
}

func (d *DefaultRefreshTokenIssuer) Configure(strategy interface{}, config sdk.Config, args ...interface{}) {
	d.RefreshTokenStrategy = strategy.(sdk.IRefreshTokenStrategy)
	for _, arg := range args {
		if ts, ok := arg.(sdk.ITokenStore); ok {
			d.TokenStore = ts
		}
	}
	if d.TokenStore == nil {
		panic("failed to init DefaultAccessTokenIssuer")
	}
	d.Lifespan = config.RefreshTokenLifespan
}
