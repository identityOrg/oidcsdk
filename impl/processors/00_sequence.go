package processors

func NewProcessorSequence(
	arg1 *DefaultBearerUserAuthProcessor,
	arg2 *DefaultClientAuthenticationProcessor,
	arg3 *DefaultGrantTypeValidator,
	arg4 *DefaultResponseTypeValidator,
	arg5 *DefaultAccessCodeValidator,
	arg6 *DefaultRefreshTokenValidator,
	arg7 *DefaultStateValidator,
	arg8 *DefaultPKCEValidator,
	arg9 *DefaultRedirectURIValidator,
	arg10 *DefaultAudienceValidationProcessor,
	arg11 *DefaultScopeValidator,
	arg12 *DefaultUserValidator,
	arg13 *DefaultClaimProcessor,
	arg14 *DefaultTokenIntrospectionProcessor,
	arg15 *DefaultTokenRevocationProcessor,
	arg16 *DefaultAuthCodeIssuer,
	arg17 *DefaultAccessTokenIssuer,
	arg18 *DefaultIDTokenIssuer,
	arg19 *DefaultRefreshTokenIssuer,
	arg20 *DefaultTokenPersister,
) []interface{} {
	seq := make([]interface{}, 0)
	seq = append(seq, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	seq = append(seq, arg8, arg9, arg10, arg11, arg12, arg13)
	seq = append(seq, arg14, arg15, arg16, arg17, arg18, arg19, arg20)
	return seq
}
