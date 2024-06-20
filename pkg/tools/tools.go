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

var opts Options

func Init(options Options) {
	opts = options
}

// SHA256 generates SHA-256 hash string.
func SHA256(text string) string {
	h := sha256.Sum256([]byte(text))
	return string(h[:])
}

// SetCookie sets cookies.
func SetCookie(c echo.Context, name, value, path string, expires time.Time) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		Path:     path,
		Expires:  expires,
		Domain:   opts.CookieDomain,
		Secure:   opts.CookieSecure,
		HttpOnly: opts.CookieHttpOnly,
	})
}
