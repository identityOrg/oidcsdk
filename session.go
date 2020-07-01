package oauth2_oidc_sdk

import (
	"net/http"
	"time"
)

type (
	ISession interface {
		GetUsername() string
		GetLoginTime() *time.Time
		IsConsentSubmitted() bool
		IsLoginDone() bool
		GetApprovedScopes() Arguments
	}

	ISessionManager interface {
		RetrieveUserSession(r *http.Request) (ISession, error)
	}
)
