package oauth2_oidc_sdk

const (
	GrantAuthorizationCode     = "authorization_code"
	GrantImplicit              = "implicit"
	GrantResourceOwnerPassword = "password"
	GrantClientCredentials     = "client_credentials"
	GrantRefreshToken          = "refresh_token"
)

const (
	ScopeOpenid        = "openid"
	ScopeProfile       = "profile"
	ScopeEmail         = "email"
	ScopeAddress       = "address"
	ScopeOfflineAccess = "offline_access"
)

const (
	ResponseTypeCode    = "code"
	ResponseTypeToken   = "token"
	ResponseTypeIdToken = "id_token"
)

const (
	ResponseModeQuery    = "query"
	ResponseModeFragment = "fragment"
	ResponseModeFormPost = "form"
)

const (
	ContentTypeUrlEncodedForm = "application/x-www-form-urlencoded"
	ContentTypeJson           = "application/json"
	ContentTypeHtml           = "text/html"
)

const (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
)
