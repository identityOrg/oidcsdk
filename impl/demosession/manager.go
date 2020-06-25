package demosession

import (
	"github.com/gorilla/sessions"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"time"
)

var sessionStore = sessions.NewCookieStore([]byte("demo-session-key"))

type Manager struct {
}

func (m *Manager) RetrieveUserSession(r *http.Request) (sess sdk.ISession, err error) {
	sessBack, err := sessionStore.Get(r, "oauth-sdk")
	if err != nil {
		return
	}
	sess = &DefaultSession{
		Username:  sessBack.Values["username"].(string),
		LoginTime: sessBack.Values["login-time"].(*time.Time),
	}
	return
}

type DefaultSession struct {
	Username  string
	LoginTime *time.Time
}

func (d DefaultSession) GetUsername() string {
	return d.Username
}

func (d DefaultSession) GetLoginTime() *time.Time {
	return d.LoginTime
}
