package secretkey

import (
	"context"
	"github.com/google/uuid"
	"github.com/identityOrg/oidcsdk/util"
	"gopkg.in/square/go-jose.v2"
)

type DefaultMemorySecretStore struct {
	Keys *jose.JSONWebKeySet
}

func (d *DefaultMemorySecretStore) GetAllSecrets(ctx context.Context) (*jose.JSONWebKeySet, error) {
	return d.Keys, nil
}

func NewDefaultMemorySecretStore() *DefaultMemorySecretStore {
	private, _ := util.GenerateRSAKeyPair(2048)
	key := jose.JSONWebKey{
		Key:       private,
		KeyID:     uuid.New().String(),
		Algorithm: "RS256",
		Use:       "sig",
	}
	return &DefaultMemorySecretStore{
		Keys: &jose.JSONWebKeySet{
			Keys: []jose.JSONWebKey{key},
		},
	}
}
