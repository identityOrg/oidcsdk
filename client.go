package oidcsdk

import "gopkg.in/square/go-jose.v2"

type IClient interface {
	GetID() string
	GetSecret() string
	IsPublic() bool
	GetIDTokenSigningAlg() jose.SignatureAlgorithm
	GetRedirectURIs() []string
	GetPostLogoutRedirectURIs() []string
	GetApprovedScopes() Arguments
	GetApprovedGrantTypes() Arguments
}
