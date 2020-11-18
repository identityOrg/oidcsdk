package oidcsdk

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
		GetScope() string
		Logout()
		Save() error
	}

	ISessionManager interface {
		RetrieveUserSession(w http.ResponseWriter, r *http.Request) (ISession, error)
	}
)
