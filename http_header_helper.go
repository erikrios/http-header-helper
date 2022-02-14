package httpheaderhelper

import (
	"net/http"
	"strings"
)

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
