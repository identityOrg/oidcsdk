package oauth2_oidc_sdk

import (
	"net/http"
	"time"
)

type (
	ISession interface {
		GetUsername() string
		GetLoginTime() time.Time
	}

	SessionFactory func(r *http.Request) (session ISession)
)
