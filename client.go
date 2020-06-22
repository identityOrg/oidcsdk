package oauth2_oidc_sdk

import "gopkg.in/square/go-jose.v2"

type IClient interface {
	GetID() string
	GetSecret() string
	IsPublic() bool
	GetIDTokenSigningAlg() jose.SignatureAlgorithm
	GetRedirectURIs() []string
	GetApprovedScopes() Arguments
	GetApprovedGrantTypes() Arguments
}
