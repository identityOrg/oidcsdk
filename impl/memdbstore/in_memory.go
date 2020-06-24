package memdbstore

import (
	"errors"
	"github.com/hashicorp/go-memdb"
	sdk "oauth2-oidc-sdk"
	client2 "oauth2-oidc-sdk/impl/client"
	"oauth2-oidc-sdk/impl/sdkerror"
	"oauth2-oidc-sdk/impl/userprofile"
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
			ID:                 "client",
			Secret:             "client",
			Public:             false,
			IDTokenSigningAlg:  "as",
			RedirectURIs:       []string{"http://localhost:8080/redirect"},
			ApprovedScopes:     []string{"openid"},
			ApprovedGrantTypes: []string{"authorization_grant", "implicit", "password", "refresh_token", "client_credentials"},
		}
		i.demo = make(map[string]interface{})
		i.demo["client"] = &client
	}
	return i
}

func (i *InMemoryDB) GetClient(clientID string) (client sdk.IClient, err error) {
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

func (i *InMemoryDB) Authenticate(username string, credential []byte) (err error) {
	if username == string(credential) {
		return nil
	} else {
		return sdkerror.InvalidGrant.WithDescription("invalid user credentials")
	}
}

func (i *InMemoryDB) GetClaims(username string, scopes sdk.Arguments, claimsIDs []string) (map[string]interface{}, error) {
	panic("implement me")
}

func (i *InMemoryDB) IsConsentRequired(username string, client sdk.IClient, scopes sdk.Arguments, audience sdk.Arguments) bool {
	panic("implement me")
}

func (i *InMemoryDB) FetchUserProfile(username string) sdk.IProfile {
	return &userprofile.DefaultProfile{
		Username: username,
	}
}

func (i *InMemoryDB) FetchClientProfile(username string) sdk.IProfile {
	return &userprofile.DefaultProfile{
		Username: username,
	}
}

func (i *InMemoryDB) Configure(interface{}, *sdk.Config, ...interface{}) {
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
}

func (i *InMemoryDB) StoreTokenProfile(reqId string, signatures sdk.TokenSignatures, profile sdk.IProfile) (err error) {
	row := &TokenTable{
		RequestID:       reqId,
		TokenSignatures: signatures,
		Profile:         profile,
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

func (i *InMemoryDB) GetProfileWithAuthCodeSign(signature string) (profile sdk.IProfile, reqId string, err error) {
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

func (i *InMemoryDB) GetProfileWithAccessTokenSign(signature string) (profile sdk.IProfile, reqId string, err error) {
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

func (i *InMemoryDB) GetProfileWithRefreshTokenSign(signature string) (profile sdk.IProfile, reqId string, err error) {
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

func (i *InMemoryDB) InvalidateWithRequestID(reqID string, what uint8) (err error) {
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
		Profile sdk.IProfile
	}
)
