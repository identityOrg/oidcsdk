package processors

func CreateDefaultSequence() []interface{} {
	var sequence []interface{}

	sequence = append(sequence, DefaultClientAuthenticationProcessor{})
	sequence = append(sequence, DefaultGrantTypeValidator{})
	sequence = append(sequence, DefaultResponseTypeValidator{})
	sequence = append(sequence, DefaultAccessCodeValidator{})
	sequence = append(sequence, DefaultRefreshTokenValidator{})
	sequence = append(sequence, DefaultRedirectURIValidator{})
	sequence = append(sequence, DefaultScopeValidator{})

	return sequence
}
