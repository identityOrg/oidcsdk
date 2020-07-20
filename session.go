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
	}

	ISessionManager interface {
		RetrieveUserSession(r *http.Request) (ISession, error)
	}
)
