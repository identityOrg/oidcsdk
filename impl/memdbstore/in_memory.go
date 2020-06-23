package memdbstore

import (
	"errors"
	"github.com/hashicorp/go-memdb"
	sdk "oauth2-oidc-sdk"
	client2 "oauth2-oidc-sdk/impl/client"
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
	panic("implement me")
}

func (i *InMemoryDB) GetClaims(username string, scopes sdk.Arguments, claimsIDs []string) (map[string]interface{}, error) {
	panic("implement me")
}

func (i *InMemoryDB) IsConsentRequired(username string, client sdk.IClient, scopes sdk.Arguments, audience sdk.Arguments) bool {
	panic("implement me")
}

func (i *InMemoryDB) Configure(interface{}, sdk.Config, ...interface{}) {
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
	row := TokeTable{
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

func (i *InMemoryDB) GetProfileWithAuthCodeSign(signature string) (profile sdk.IProfile, err error) {
	panic("implement me")
}

func (i *InMemoryDB) GetProfileWithAccessTokenSign(signature string) (profile sdk.IProfile, err error) {
	panic("implement me")
}

func (i *InMemoryDB) GetProfileWithRefreshTokenSign(signature string) (profile sdk.IProfile, err error) {
	panic("implement me")
}

func (i *InMemoryDB) InvalidateWithRequestID(reqID string) (err error) {
	panic("implement me")
}

type (
	TokeTable struct {
		RequestID string
		sdk.TokenSignatures
		Profile sdk.IProfile
	}
)
