package secretkey

import (
	"github.com/google/uuid"
	"github.com/identityOrg/oidcsdk/util"
	"gopkg.in/square/go-jose.v2"
)

type DefaultMemorySecretStore struct {
	Keys *jose.JSONWebKeySet
}

func NewDefaultMemorySecretStore() *DefaultMemorySecretStore {
	private, _ := util.GenerateRSAKeyPair(2048)
	key := jose.JSONWebKey{
		Key:       private,
		KeyID:     uuid.New().String(),
		Algorithm: "RS256",
		Use:       "sign",
	}
	return &DefaultMemorySecretStore{
		Keys: &jose.JSONWebKeySet{
			Keys: []jose.JSONWebKey{key},
		},
	}
}

func (d *DefaultMemorySecretStore) GetAllSecrets() *jose.JSONWebKeySet {
	return d.Keys
}
