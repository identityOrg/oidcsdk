package demosession

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/example/config"
	"github.com/identityOrg/oidcsdk/util"
	"net/http"
	"strings"
	"time"
)

func init() {
	gob.Register(&time.Time{})
}

const (
	ScopeAttribute        = "scope"
	LoginTimeAttribute    = "login-time"
	UsernameAttribute     = "username"
	LoginFlashAttribute   = "login"
	ConsentFlashAttribute = "consent"
)

type Manager struct {
	SessionStore *sessions.CookieStore
	CookieName   string
}

func NewManager(config *config.DemoConfig) *Manager {
	return &Manager{
		SessionStore: sessions.NewCookieStore([]byte(config.SessionEncKey)),
		CookieName:   config.SessionCookieName,
	}
}

func (m *Manager) RetrieveUserSession(w http.ResponseWriter, r *http.Request) (sdk.ISession, error) {
	sessBack, err := m.SessionStore.Get(r, "oauth-sdk")
	if err != nil {
		return nil, err
	}
	sess := &DefaultSession{
		s: sessBack,
		r: r,
		w: w,
	}
	return sess, nil
}

//func (m *Manager) StoreUserSession(w http.ResponseWriter, r *http.Request, sess sdk.ISession) error {
//	sessBack, err := m.SessionStore.Get(r, "oauth-sdk")
//	if err != nil {
//		return err
//	}
//	sessBack.Values[UsernameAttribute] = sess.GetUsername()
//	sessBack.Values[ScopeAttribute] = sess.GetScope()
//	sessBack.Values[LoginTimeAttribute] = sess.GetLoginTime().Unix()
//
//	return m.SessionStore.Save(r, w, sessBack)
//}

type DefaultSession struct {
	s *sessions.Session
	r *http.Request
	w http.ResponseWriter
}

func (d DefaultSession) SetAttribute(name string, value interface{}) {
	d.s.Values[name] = value
}

func (d DefaultSession) GetAttribute(name string) interface{} {
	return d.s.Values[name]
}

func (d DefaultSession) Save() error {
	return d.s.Save(d.r, d.w)
}

func (d DefaultSession) Logout() {
	d.s.Options.MaxAge = -1
}

func (d DefaultSession) GetScope() string {
	scope := d.s.Values[ScopeAttribute]
	if scope != nil {
		return scope.(string)
	}
	return ""
}

func (d DefaultSession) IsLoginDone() bool {
	logins := d.s.Flashes(LoginFlashAttribute)
	if len(logins) > 0 {
		return true
	}
	return false
}

func (d DefaultSession) IsConsentSubmitted() bool {
	consents := d.s.Flashes(ConsentFlashAttribute)
	if len(consents) > 0 {
		return true
	}
	return false
}

func (d DefaultSession) GetApprovedScopes() sdk.Arguments {
	return util.RemoveEmpty(strings.Split(d.GetScope(), " "))
}

func (d DefaultSession) GetUsername() string {
	userName := d.s.Values[UsernameAttribute]
	if userName != nil {
		return userName.(string)
	}
	return ""
}

func (d DefaultSession) GetLoginTime() *time.Time {
	loginTime := d.s.Values[LoginTimeAttribute]
	if loginTime != nil {
		return loginTime.(*time.Time)
	}
	return nil
}
