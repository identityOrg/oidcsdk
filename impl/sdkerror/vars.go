package sdkerror

var (
	InvalidRequest = DefaultErrorFactory(400, "invalid_request",
		"The request is missing a required parameter, includes an unsupported parameter value")
	InvalidClient = DefaultErrorFactory(401, "invalid_client",
		"Client authentication failed (e.g., unknown client, no client authentication included, or"+
			" unsupported authentication method)")
	InvalidGrant = DefaultErrorFactory(400, "invalid_grant",
		"The provided authorization grant (e.g., authorizationcode, resource owner credentials) or"+
			" refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization"+
			" request, or was issued to another client")
	UnAuthorizedClient = DefaultErrorFactory(400, "unauthorized_client",
		"The authenticated client is not authorized to use this authorization grant type")
	UnSupportedGrantType = DefaultErrorFactory(400, "unsupported_grant_type",
		"The authorization grant type is not supported by the authorization server")
	InvalidScope = DefaultErrorFactory(400, "invalid_scope", "Grant is invalid")
)
