package oauth2_oidc_sdk

import "gopkg.in/square/go-jose.v2"

type IClient interface {
	GetID() string
	GetSecret() string
	GetIDTokenSigningAlg() jose.SignatureAlgorithm
}
