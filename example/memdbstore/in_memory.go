package memdbstore

import (
	"context"
	"errors"
	"github.com/hashicorp/go-memdb"
	sdk "github.com/identityOrg/oidcsdk"
	client2 "github.com/identityOrg/oidcsdk/impl/client"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
	"gopkg.in/square/go-jose.v2"
	"time"
)

type InMemoryDB struct {
	Db   *memdb.MemDB
	demo map[string]interface{}
}

func NewInMemoryDB(demo bool) *InMemoryDB {
	i := &InMemoryDB{}
	if demo {
		client := client2.DefaultClient{
			ID:                "client",
			Secret:            "client",
			Public:            false,
			IDTokenSigningAlg: jose.RS256,
			RedirectURIs: []string{
				"http://localhost:8080/redirect",
				"http://client.localhost:4200/login/oauth2/code/goid",
			},
			ApprovedScopes:     []string{sdk.ScopeOpenid},
			ApprovedGrantTypes: []string{sdk.GrantAuthorizationCode, sdk.GrantImplicit, sdk.GrantResourceOwnerPassword, sdk.GrantRefreshToken, sdk.GrantClientCredentials},
		}
		i.demo = make(map[string]interface{})
		i.demo["client"] = &client
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic("failed to init InMemoryDB" + err.Error())
	}
	i.Db = db
	if len(i.demo) > 0 {
		txn := i.Db.Txn(true)
		for k, v := range i.demo {
			err := txn.Insert(k, v)
			if err != nil {
				txn.Abort()
				panic("failed to create demo data " + err.Error())
			}
		}
		txn.Commit()
	}
	i.Db = db
	return i
}

func (i *InMemoryDB) GetClient(ctx context.Context, clientID string) (client sdk.IClient, err error) {
	txn := i.Db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("client", "id", clientID)
	if err != nil {
		return nil, err
	} else if raw == nil {
		return nil, errors.New("client not found")
	}
	return raw.(sdk.IClient), nil
}

func (i *InMemoryDB) Authenticate(ctx context.Context, username string, credential []byte) (err error) {
	if username == string(credential) {
		return nil
	} else {
		return sdkerror.ErrRequestUnauthorized.WithDescription("invalid user credentials")
	}
}

func (i *InMemoryDB) GetClaims(_ context.Context, username string, _ sdk.Arguments, _ []string) (map[string]interface{}, error) {
	claims := make(map[string]interface{})
	claims["username"] = username
	return claims, nil
}

func (i *InMemoryDB) IsConsentRequired(context.Context, string, string, sdk.Arguments) bool {
	return false
}

func (i *InMemoryDB) StoreConsent(context.Context, string, string, sdk.Arguments) error {
	return nil
}

func (i *InMemoryDB) FetchUserProfile(ctx context.Context, username string) sdk.RequestProfile {
	profile := sdk.RequestProfile{}
	profile.SetUsername(username)
	profile.SetDomain("demo.com")
	return profile
}

func (i *InMemoryDB) FetchClientProfile(ctx context.Context, username string) sdk.RequestProfile {
	profile := sdk.RequestProfile{}
	profile.SetUsername(username)
	return profile
}

func (i *InMemoryDB) StoreTokenProfile(ctx context.Context, reqId string, signatures sdk.ITokenSignatures, profile sdk.RequestProfile) (err error) {
	row := &TokenTable{
		RequestID: reqId,
		TokenSignatures: sdk.TokenSignatures{
			AuthorizationCodeSignature: signatures.GetACSignature(),
			AccessTokenSignature:       signatures.GetATSignature(),
			RefreshTokenSignature:      signatures.GetRTSignature(),
			RefreshTokenExpiry:         signatures.GetRTExpiry(),
			AccessTokenExpiry:          signatures.GetATExpiry(),
			AuthorizationCodeExpiry:    signatures.GetACExpiry(),
		},
		Profile: profile,
	}

	txn := i.Db.Txn(true)
	err = txn.Insert("request", row)
	if err != nil {
		txn.Abort()
		return err
	}
	txn.Commit()
	return nil
}

func (i *InMemoryDB) GetProfileWithAuthCodeSign(ctx context.Context, signature string) (profile sdk.RequestProfile, reqId string, err error) {
	txn := i.Db.Txn(false)
	defer txn.Abort()

	first, err := txn.First("request", "code-sign", signature)
	if err != nil {
		return nil, "", err
	} else if first == nil {
		return nil, "", errors.New("authorization code not found")
	}
	row := first.(*TokenTable)
	now := time.Now()
	if row.AuthorizationCodeExpiry.Before(now) {
		return nil, "", errors.New("authorization code expired")
	} else {
		return row.Profile, row.RequestID, nil
	}
}

func (i *InMemoryDB) GetProfileWithAccessTokenSign(ctx context.Context, signature string) (profile sdk.RequestProfile, reqId string, err error) {
	txn := i.Db.Txn(false)
	defer txn.Abort()

	first, err := txn.First("request", "at-sign", signature)
	if err != nil {
		return nil, "", err
	} else if first == nil {
		return nil, "", errors.New("access token not found")
	}
	row := first.(*TokenTable)
	now := time.Now()
	if row.AccessTokenExpiry.Before(now) {
		return nil, "", errors.New("access token expired")
	} else {
		return row.Profile, row.RequestID, nil
	}
}

func (i *InMemoryDB) GetProfileWithRefreshTokenSign(ctx context.Context, signature string) (profile sdk.RequestProfile, reqId string, err error) {
	txn := i.Db.Txn(false)
	defer txn.Abort()

	first, err := txn.First("request", "rt-sign", signature)
	if err != nil {
		return nil, "", err
	} else if first == nil {
		return nil, "", errors.New("refresh token not found")
	}
	row := first.(*TokenTable)
	now := time.Now()
	if row.RefreshTokenExpiry.Before(now) {
		return nil, "", errors.New("refresh token expired")
	} else {
		return row.Profile, row.RequestID, nil
	}
}

func (i *InMemoryDB) InvalidateWithRequestID(ctx context.Context, reqID string, what uint8) (err error) {
	txn := i.Db.Txn(true)
	first, err := txn.First("request", "id", reqID)
	if err != nil {
		return
	}
	if first != nil {
		row := first.(*TokenTable)
		if what&sdk.ExpireRefreshToken > 0 {
			row.RefreshTokenExpiry = time.Now()
		}
		if what&sdk.ExpireAccessToken > 0 {
			row.AccessTokenExpiry = time.Now()
		}
		if what&sdk.ExpireAuthorizationCode > 0 {
			row.AuthorizationCodeExpiry = time.Now()
		}
		txn.Commit()
		return nil
	}
	txn.Abort()
	return nil
}

type (
	TokenTable struct {
		RequestID string
		sdk.TokenSignatures
		Profile sdk.RequestProfile
	}
)
