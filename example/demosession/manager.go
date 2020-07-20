package demosession

import (
	"github.com/gorilla/sessions"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/util"
	"net/http"
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
		i := loginTime.(int64)
		unix := time.Unix(i, 0)
		sess.LoginTime = &unix
	}
	if scope != nil {
		sess.Scope = scope.(string)
	}
	return sess, nil
}

func (m *Manager) StoreUserSession(w http.ResponseWriter, r *http.Request, sess sdk.ISession) error {
	sessBack, err := m.SessionStore.Get(r, "oauth-sdk")
	if err != nil {
		return err
	}
	sessBack.Values["username"] = sess.GetUsername()
	sessBack.Values["scope"] = sess.GetScope()
	sessBack.Values["login-time"] = sess.GetLoginTime().Unix()

	return m.SessionStore.Save(r, w, sessBack)
}

type DefaultSession struct {
	Username  string
	Scope     string
	LoginTime *time.Time
}

func (d DefaultSession) GetScope() string {
	return d.Scope
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
