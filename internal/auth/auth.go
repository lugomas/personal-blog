package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

var (
	// Define the expected admin username and password
	adminUsername = "admin"
	adminPassword = "password123"
)

// BasicAuthMiddleware is a middleware function that enforces basic authentication.
func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the "Authorization" header from the incoming HTTP request
		auth := r.Header.Get("Authorization")

		// If the "Authorization" header is missing, prompt the user for credentials
		if auth == "" {
			// Set the "WWW-Authenticate" header to indicate that basic authentication is required
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			// Respond with a 401 Unauthorized error
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode the base64-encoded credentials from the "Authorization" header
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			// If decoding fails, respond with a 401 Unauthorized error
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Split the decoded string into "username:password" format
		pair := strings.SplitN(string(payload), ":", 2)

		// If the split fails or the credentials do not match the expected values, reject the request
		if len(pair) != 2 || pair[0] != adminUsername || pair[1] != adminPassword {
			// Respond with a 401 Unauthorized error if the credentials are incorrect
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the credentials are valid, continue processing the request by calling the next handler
		next.ServeHTTP(w, r)
	}
}
