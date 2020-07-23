package oidcsdk

import (
	"context"
	"gopkg.in/square/go-jose.v2"
	"time"
)

type (
	ITokenStore interface {
		StoreTokenProfile(ctx context.Context, reqId string, signatures TokenSignatures, profile RequestProfile) (err error)
		GetProfileWithAuthCodeSign(ctx context.Context, signature string) (profile RequestProfile, reqId string, err error)
		GetProfileWithAccessTokenSign(ctx context.Context, signature string) (profile RequestProfile, reqId string, err error)
		GetProfileWithRefreshTokenSign(ctx context.Context, signature string) (profile RequestProfile, reqId string, err error)
		InvalidateWithRequestID(ctx context.Context, reqID string, what uint8) (err error)
	}

	ITokenStoreNew interface {
		StoreAuthorizationCode(ctx context.Context, reqId string, sign string, expiry time.Time, profile RequestProfile) (err error)
		StoreAccessToken(ctx context.Context, reqId string, sign string, expiry time.Time, profile RequestProfile) (err error)
		StoreRefreshToken(ctx context.Context, reqId string, sign string, expiry time.Time, profile RequestProfile) (err error)
		InvalidateAuthorizationCode(ctx context.Context, reqId string) (err error)
		InvalidateAccessToken(ctx context.Context, reqId string) (err error)
		InvalidateRefreshToken(ctx context.Context, reqId string) (err error)
		FindAuthorizationCode(ctx context.Context, sign string)
	}

	IUserStore interface {
		Authenticate(ctx context.Context, username string, credential []byte) (err error)
		GetClaims(ctx context.Context, username string, scopes Arguments, claimsIDs []string) (map[string]interface{}, error)
		IsConsentRequired(ctx context.Context, username string, clientId string, scopes Arguments) bool
		StoreConsent(ctx context.Context, username string, clientId string, scopes Arguments) error
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

	ISecretStore interface {
		GetAllSecrets() *jose.JSONWebKeySet
	}
)

const (
	ExpireAuthorizationCode = 1
	ExpireAccessToken       = 2
	ExpireRefreshToken      = 4
)
