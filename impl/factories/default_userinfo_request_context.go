package factories

import sdk "github.com/identityOrg/oidcsdk"

type DefaultUserInfoRequestContext struct {
	BearerToken     string
	Username        string
	Claims          map[string]interface{}
	ApprovedScopes  sdk.Arguments
	RequestedClaims []string
}

func (d *DefaultUserInfoRequestContext) GetApprovedScopes() sdk.Arguments {
	return d.ApprovedScopes
}

func (d *DefaultUserInfoRequestContext) SetApprovedScopes(scopes sdk.Arguments) {
	d.ApprovedScopes = scopes
}

func (d *DefaultUserInfoRequestContext) GetRequestedClaims() []string {
	return d.RequestedClaims
}

func (d *DefaultUserInfoRequestContext) SetRequestedClaims(claimIds []string) {
	d.RequestedClaims = claimIds
}

func (d *DefaultUserInfoRequestContext) GetBearerToken() string {
	return d.BearerToken
}

func (d *DefaultUserInfoRequestContext) GetUsername() string {
	return d.Username
}

func (d *DefaultUserInfoRequestContext) SetUsername(username string) {
	d.Username = username
}

func (d *DefaultUserInfoRequestContext) GetClaims() map[string]interface{} {
	return d.Claims
}

func (d *DefaultUserInfoRequestContext) AddClaim(claimId string, value interface{}) {
	d.Claims[claimId] = value
}
