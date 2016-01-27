package auth

import "net/http"

// Info is the info that used for check authentification
type Info struct {
	User     string
	Password string
	Realm    string
}

// Wrapper is the object to wrap http.HandlerFunc and provide auth
type Wrapper interface {
	Wrap(wrapped http.HandlerFunc) http.HandlerFunc
}

// NewBaseAuth create new base authentification
func NewBaseAuth(info Info) Wrapper {
	return base{info}
}
