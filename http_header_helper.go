package httpheaderhelper

import (
	"net/http"
	"strings"
)

// IsContentTypeApplicationJson is a function the content type of HTTP request header.
// If the content type is application/json, return true.
// Otherwise, return false.
func IsContentTypeApplicationJson(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

// IsContainsBasicAuthenctication check the authentication schema.
// Return true if the schema is valid.
// Otherwise, return false.
func IsContainsBasicAuthenctication(r *http.Request) bool {
	stringValue := r.Header.Get("Authorization")
	values := strings.Split(stringValue, " ")

	if len(values) != 2 {
		return false
	}

	if values[0] != "Basic" {
		return false
	}

	return true
}
