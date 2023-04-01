package sessions

import (
	"banana/pkg/domain"

	"net/http"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const sessionName = "banana-session"

var store = sessions.NewFilesystemStore(os.TempDir(), securecookie.GenerateRandomKey(32))

func StartSession(w http.ResponseWriter, r *http.Request, id uint64) error {
	session, _ := store.Get(r, sessionName)
	session.Values["id"] = id
	session.Options = &sessions.Options{
		MaxAge:   100000,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
	}
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func FinishSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}

	_, isIn := session.Values["id"]
	if !isIn {
		return domain.ErrFinishSession
	}

	session.Options.MaxAge = -1
	session.Options.Secure = true
	session.Options.SameSite = http.SameSiteNoneMode

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func CheckSession(r *http.Request) (uint64, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return 0, err
	}

	id, isIn := session.Values["id"]
	if !isIn || session.IsNew {
		return 0, domain.ErrUserNotLoggedIn
	}

	idCasted, ok := id.(uint64)
	if !ok {
		return 0, domain.ErrSessionCast
	}

	return idCasted, nil
}
