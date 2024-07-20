package tools

import (
	"crypto/sha256"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Options struct {
	CookieDomain   string
	CookieSecure   bool
	CookieHttpOnly bool
}

var options Options

func Init(opts Options) {
	options = opts
}

// SHA256 generates SHA-256 hash string.
func SHA256(text string) string {
	h := sha256.Sum256([]byte(text))
	return string(h[:])
}

// SetCookie sets cookies.
func SetCookie(c echo.Context, name, path, value string, expires time.Time) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Path:     path,
		Value:    value,
		Expires:  expires,
		Domain:   options.CookieDomain,
		Secure:   options.CookieSecure,
		HttpOnly: options.CookieHttpOnly,
	})
}
