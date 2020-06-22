package compose

import "oauth2-oidc-sdk/impl/processors"

func CreateDefaultSequence() []interface{} {
	var sequence []interface{}

	sequence = append(sequence, &processors.DefaultClientAuthenticationProcessor{})
	sequence = append(sequence, &processors.DefaultGrantTypeValidator{})
	sequence = append(sequence, &processors.DefaultResponseTypeValidator{})
	sequence = append(sequence, &processors.DefaultAccessCodeValidator{})
	sequence = append(sequence, &processors.DefaultRefreshTokenValidator{})
	sequence = append(sequence, &processors.DefaultRedirectURIValidator{})
	sequence = append(sequence, &processors.DefaultScopeValidator{})

	return sequence
}
