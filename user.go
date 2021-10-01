package oidcsdk

import (
	"github.com/identityOrg/oidcsdk/util"
	"strings"
)

type RequestProfile map[string]string

func NewRequestProfile() RequestProfile {
	return make(map[string]string)
}

func (r RequestProfile) GetUsername() string {
	return r["username"]
}

func (r RequestProfile) SetUsername(username string) {
	r["username"] = username
}

func (r RequestProfile) GetClientID() string {
	return r["client_id"]
}

func (r RequestProfile) SetClientID(username string) {
	r["client_id"] = username
}

func (r RequestProfile) GetState() string {
	return r["state"]
}

func (r RequestProfile) SetState(state string) {
	r["state"] = state
}

func (r RequestProfile) GetNonce() string {
	return r["nonce"]
}

func (r RequestProfile) SetNonce(nonce string) {
	r["nonce"] = nonce
}

func (r RequestProfile) GetRedirectURI() string {
	return r["redirect_uri"]
}

func (r RequestProfile) SetRedirectURI(redirectUri string) {
	r["redirect_uri"] = redirectUri
}

func (r RequestProfile) GetScope() Arguments {
	s := r["scope"]
	if s != "" {
		return util.RemoveEmpty(strings.Split(s, " "))
	}
	return []string{}
}

func (r RequestProfile) SetScope(scopes Arguments) {
	r["scope"] = strings.Join(scopes, " ")
}

func (r RequestProfile) GetAudience() Arguments {
	s := r["audience"]
	if s != "" {
		return util.RemoveEmpty(strings.Split(s, " "))
	}
	return []string{}
}

func (r RequestProfile) SetAudience(aud Arguments) {
	r["audience"] = strings.Join(aud, " ")
}

func (r RequestProfile) IsClient() bool {
	return r["domain"] == ""
}

func (r RequestProfile) GetDomain() string {
	return r["domain"]
}

func (r RequestProfile) SetDomain(domain string) {
	r["domain"] = domain
}

func (r RequestProfile) GetCodeChallenge() string {
	return r["code_challenge"]
}

func (r RequestProfile) SetCodeChallenge(challenge string) {
	r["code_challenge"] = challenge
}

func (r RequestProfile) GetCodeChallengeMethod() string {
	return r["code_challenge_method"]
}

func (r RequestProfile) SetCodeChallengeMethod(challengeMethod string) {
	r["code_challenge_method"] = challengeMethod
}
func (r RequestProfile) GetGrantType() string {
	return r["grant_type"]
}

func (r RequestProfile) SetGrantType(challengeMethod string) {
	r["grant_type"] = challengeMethod
}
