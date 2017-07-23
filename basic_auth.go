package basicauth

import (
	"net/http"
)

// With wraps the passed handler in BasicAuth
func With(auth Credentials) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if auth.isAuthenticated(r) {
				next.ServeHTTP(w, r)
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted Area"`)
			http.Error(w, "Not authorized", http.StatusUnauthorized)
		})
	}
}

// Credentials describes the user password tuple used by the middleware
type Credentials struct {
	User string
	Pass string
}

func (c *Credentials) skip() bool {
	return c.Pass == ""
}

func (c *Credentials) isAuthenticated(r *http.Request) bool {
	if c.skip() {
		return true
	}
	if user, pass, ok := r.BasicAuth(); ok {
		return user == c.User && pass == c.Pass
	}
	return false
}
