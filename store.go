package oauth2_oidc_sdk

import "context"

type (
	ITokenStore interface {
		StoreTokenProfile(regID string, signatures TokenSignatures, profile IProfile) (err error)
		GetProfileWithAuthCodeSign(signature string) (profile IProfile, reqId string, err error)
		GetProfileWithAccessTokenSign(signature string) (profile IProfile, reqId string, err error)
		GetProfileWithRefreshTokenSign(signature string) (profile IProfile, reqId string, err error)
		InvalidateWithRequestID(reqID string, what uint8) (err error)
	}

	IUserStore interface {
		Authenticate(username string, credential []byte) (err error)
		GetClaims(username string, scopes Arguments, claimsIDs []string) (map[string]interface{}, error)
		IsConsentRequired(username string, client IClient, scopes Arguments, audience Arguments) bool
		FetchUserProfile(username string) IProfile
	}

	IClientStore interface {
		GetClient(clientID string) (client IClient, err error)
		FetchClientProfile(clientID string) IProfile
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
