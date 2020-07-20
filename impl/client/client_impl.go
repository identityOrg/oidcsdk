package client

import (
	sdk "github.com/identityOrg/oidcsdk"
	"gopkg.in/square/go-jose.v2"
)

type DefaultClient struct {
	ID                 string
	Secret             string
	Public             bool
	IDTokenSigningAlg  jose.SignatureAlgorithm
	RedirectURIs       []string
	ApprovedScopes     sdk.Arguments
	ApprovedGrantTypes sdk.Arguments
}

func (d *DefaultClient) GetID() string {
	return d.ID
}

func (d *DefaultClient) GetSecret() string {
	return d.Secret
}

func (d *DefaultClient) IsPublic() bool {
	return d.Public
}

func (d *DefaultClient) GetIDTokenSigningAlg() jose.SignatureAlgorithm {
	return d.IDTokenSigningAlg
}

func (d *DefaultClient) GetRedirectURIs() []string {
	return d.RedirectURIs
}

func (d *DefaultClient) GetApprovedScopes() sdk.Arguments {
	return d.ApprovedScopes
}

func (d *DefaultClient) GetApprovedGrantTypes() sdk.Arguments {
	return d.ApprovedGrantTypes
}
