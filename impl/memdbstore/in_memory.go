package memdbstore

import (
	"github.com/hashicorp/go-memdb"
	sdk "oauth2-oidc-sdk"
)

type InMemoryDB struct {
	Db *memdb.MemDB
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{}
}

func (i *InMemoryDB) GetClient(clientID string) (client sdk.IClient, err error) {
	panic("implement me")
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
