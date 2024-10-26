package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("API key not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Invalid authorization type")
	}

	return vals[1], nil
}
