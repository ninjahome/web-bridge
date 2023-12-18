package server

import (
	"github.com/gorilla/sessions"
	"net/http"
	"sync"
)

var _instance *SessionManager
var logOnce sync.Once

const SessionNameTwitter = "session-name-for-twitter"

type SessionManager struct {
	store *sessions.CookieStore
}

func SMInst() *SessionManager {
	logOnce.Do(func() {
		_instance = newSM()
	})
	return _instance
}

func newSM() *SessionManager {
	var store = sessions.NewCookieStore([]byte(_globalCfg.SessionKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	return &SessionManager{
		store: store,
	}
}

func (sm *SessionManager) Get(key string, r *http.Request) (any, error) {
	session, err := sm.store.Get(r, SessionNameTwitter)
	if err != nil {
		return nil, err
	}

	data := session.Values[key]
	return data, nil
}

func (sm *SessionManager) Set(r *http.Request, w http.ResponseWriter, key string, val any) error {
	session, err := sm.store.Get(r, SessionNameTwitter)
	if err != nil {
		return err
	}

	session.Values[key] = val
	return session.Save(r, w)
}
