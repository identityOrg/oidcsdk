package oauth2_oidc_sdk

type (
	ITokenStore interface {
		StoreTokenProfile(regID string, signatures ITokenSignatures, profile IProfile) (err error)
		GetProfileWithAuthCodeSign(signature string) (profile IProfile, err error)
		GetProfileWithAccessTokenSign(signature string) (profile IProfile, err error)
		GetProfileWithRefreshTokenSign(signature string) (profile IProfile, err error)
		InvalidateWithRequestID(reqID string) (err error)
	}

	IUserStore interface {
		Authenticate(username string, credential []byte) (err error)
		GetClaims(username string, claimsIDs []string) (map[string]interface{}, error)
		IsConsentRequired(username string, client IClient, scopes Arguments, audience Arguments) bool
	}

	IClientStore interface {
		GetClient(clientID string) (client IClient, err error)
	}
)
