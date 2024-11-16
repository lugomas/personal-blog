package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

var (
	adminUsername = "admin"
	adminPassword = "password123"
)

func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode base64 encoded "username:password"
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		// Split username and password
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || pair[0] != adminUsername || pair[1] != adminPassword {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Continue to the handler if credentials are valid
		next.ServeHTTP(w, r)
	}
}
