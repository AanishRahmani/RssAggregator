package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract an apikey from the http request
// eg: authorization:apikey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {

	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth error")
	}
	return vals[1], nil
}
