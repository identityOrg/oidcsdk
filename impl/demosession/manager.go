package demosession

import (
	"github.com/gorilla/sessions"
	"net/http"
	sdk "oauth2-oidc-sdk"
	"oauth2-oidc-sdk/util"
	"strings"
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
	scope := sessBack.Values["scope"]
	loginTime := sessBack.Values["login-time"]
	if userName != nil {
		sess.Username = userName.(string)
	}
	if loginTime != nil {
		sess.LoginTime = loginTime.(*time.Time)
	}
	if scope != nil {
		sess.Scope = scope.(string)
	}
	return sess, nil
}

type DefaultSession struct {
	Username  string
	Scope     string
	LoginTime *time.Time
}

func (d DefaultSession) IsLoginDone() bool {
	return true
}

func (d DefaultSession) IsConsentSubmitted() bool {
	return true
}

func (d DefaultSession) GetApprovedScopes() sdk.Arguments {
	return util.RemoveEmpty(strings.Split(d.Scope, " "))
}

func (d DefaultSession) GetUsername() string {
	return d.Username
}

func (d DefaultSession) GetLoginTime() *time.Time {
	return d.LoginTime
}
