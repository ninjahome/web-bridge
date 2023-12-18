package server

import (
	"github.com/gorilla/sessions"
	"github.com/ninjahome/web-bridge/util"
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
	secretKey, err := util.RandomBytesInHex(32)
	if err != nil {
		panic(err)
	}
	var store = sessions.NewCookieStore([]byte(secretKey))
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

	// 从会话中获取值
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
