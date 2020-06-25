package demosession

import (
	"github.com/gorilla/sessions"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"time"
)

type Manager struct {
	SessionStore *sessions.CookieStore
	CookieName   string
}

func NewManager(encKey string, cookieName string) *Manager {
	return &Manager{
		SessionStore: sessions.NewCookieStore([]byte(encKey)),
		CookieName:   cookieName,
	}
}

func (m *Manager) RetrieveUserSession(r *http.Request) (sdk.ISession, error) {
	sessBack, err := m.SessionStore.Get(r, "oauth-sdk")
	if err != nil {
		return nil, err
	}
	sess := &DefaultSession{}
	userName := sessBack.Values["username"]
	loginTime := sessBack.Values["login-time"]
	if userName != nil {
		sess.Username = userName.(string)
	}
	if loginTime != nil {
		sess.LoginTime = loginTime.(*time.Time)
	}
	return sess, nil
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
