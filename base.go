package auth

import "net/http"

type base struct {
	info Info
}

func (base base) Wrap(wrapped http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if base.isOk(r) {
			wrapped(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=\""+base.info.Realm+"\"")
			http.Error(w, "Unauthorized", 401)
		}
	}
}

func (base base) isOk(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	if username != base.info.User {
		return false
	}
	return password == base.info.Password
}
