package util

import "net/http"

// IsHTMX checks if the HTTP request is made using HTMX.
// HTMX requests typically include a specific header "HX-Request".
// This function checks for the presence of this header and
// returns true if it exists, indicating an HTMX request.
//
// Usage:
//
//	if util.IsHTMX(r) {
//	    // Handle HTMX specific logic
//	}
//
// Params:
//
//	r *http.Request - The HTTP request to check.
//
// Returns:
//
//	bool - True if the request is an HTMX request, false otherwise.
func IsHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") != ""
}

// RespondWithError is a generic utility function for sending error responses.
func RespondWithError(w http.ResponseWriter, errorMsg string, statusCode int) {
	http.Error(w, errorMsg, statusCode)
}
