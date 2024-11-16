package handler

import "net/http"

// LogoutHandler handles logout requests and prompts the browser to re-authenticate.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Set the WWW-Authenticate header to prompt for credentials again
	w.Header().Set("WWW-Authenticate", `Basic realm="Authorization required"`)
	// Return a 401 Unauthorized status to trigger the browser's authentication prompt
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
