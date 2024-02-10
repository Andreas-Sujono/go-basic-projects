package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeader = errors.New("No auth header")
var ErrMalformedAuthHeader = errors.New("Maformed auth header")

func GetApiKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", ErrNoAuthHeader
	}

	splitAuth := strings.Split(auth, " ")
	if splitAuth[0] != "ApiKey" {
		return "", ErrMalformedAuthHeader
	}

	return splitAuth[1], nil
}
