package validators

import (
	"errors"
	"net/url"
	"strings"

	"github.com/canyouhearthemusic/proxy-server/internal/models"
)

var validMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}

func ValidateRequest(proxyReq *models.ProxyRequest) error {
	if !isValidMethod(proxyReq.Method) {
		return errors.New("invalid HTTP method")
	}
	if _, err := url.ParseRequestURI(proxyReq.URL); err != nil {
		return errors.New("invalid URL")
	}
	return nil
}

func isValidMethod(method string) bool {
	method = strings.ToUpper(method)
	for _, validMethod := range validMethods {
		if method == validMethod {
			return true
		}
	}

	return false
}
