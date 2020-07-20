package compose

import "github.com/identityOrg/oidcsdk/impl/processors"

func CreateDefaultSequence() []interface{} {
	var sequence []interface{}

	sequence = append(sequence, &processors.DefaultClientAuthenticationProcessor{})
	sequence = append(sequence, &processors.DefaultGrantTypeValidator{})
	sequence = append(sequence, &processors.DefaultResponseTypeValidator{})
	sequence = append(sequence, &processors.DefaultAccessCodeValidator{})
	sequence = append(sequence, &processors.DefaultRefreshTokenValidator{})
	sequence = append(sequence, &processors.DefaultRedirectURIValidator{})
	sequence = append(sequence, &processors.DefaultScopeValidator{})
	sequence = append(sequence, &processors.DefaultUserValidator{})
	sequence = append(sequence, &processors.DefaultClaimProcessor{})
	sequence = append(sequence, &processors.DefaultAuthCodeIssuer{})
	sequence = append(sequence, &processors.DefaultAccessTokenIssuer{})
	sequence = append(sequence, &processors.DefaultRefreshTokenIssuer{})
	sequence = append(sequence, &processors.DefaultIDTokenIssuer{})
	sequence = append(sequence, &processors.DefaultTokenPersister{})

	return sequence
}
