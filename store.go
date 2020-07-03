package oauth2_oidc_sdk

import "context"

type (
	ITokenStore interface {
		StoreTokenProfile(ctx context.Context, reqId string, signatures TokenSignatures, profile RequestProfile) (err error)
		GetProfileWithAuthCodeSign(ctx context.Context, signature string) (profile RequestProfile, reqId string, err error)
		GetProfileWithAccessTokenSign(ctx context.Context, signature string) (profile RequestProfile, reqId string, err error)
		GetProfileWithRefreshTokenSign(ctx context.Context, signature string) (profile RequestProfile, reqId string, err error)
		InvalidateWithRequestID(ctx context.Context, reqID string, what uint8) (err error)
	}

	IUserStore interface {
		Authenticate(ctx context.Context, username string, credential []byte) (err error)
		GetClaims(ctx context.Context, username string, scopes Arguments, claimsIDs []string) (map[string]interface{}, error)
		IsConsentRequired(ctx context.Context, username string, clientId string, scopes Arguments) bool
		FetchUserProfile(ctx context.Context, username string) RequestProfile
	}

	IClientStore interface {
		GetClient(ctx context.Context, clientID string) (client IClient, err error)
		FetchClientProfile(ctx context.Context, clientID string) RequestProfile
	}

	ITransactionalStore interface {
		StartTransaction(ctx context.Context)
		CommitTransaction(ctx context.Context)
		RollbackTransaction(ctx context.Context)
	}
)

const (
	ExpireAuthorizationCode = 1
	ExpireAccessToken       = 2
	ExpireRefreshToken      = 4
)
