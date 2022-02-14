package httpheaderhelper

import "net/http"

// IsContentTypeApplicationJson is a function the content type of HTTP request header.
// If the content type is application/json, return true.
// Otherwise, return false.
func IsContentTypeApplicationJson(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}
