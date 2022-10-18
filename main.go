// Package cookiefilter traefik middleware plugin.
package cookiefilter

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	// Specify a list of cookie names that should be preserved
	// all other cookies will be removed, an empty slice
	// that all cookies will be removed
	KeepCookies []string `json:"keepCookies,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		KeepCookies: []string{},
	}
}

type CookieFilter struct {
	next     http.Handler
	keepCookies  []string
	name     string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &CookieFilter{
		keepCookies:  config.KeepCookies,
		next:     next,
		name:     name,
	}, nil
}

func (c *CookieFilter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Copy original request
	originalRequest := req.Clone(req.Context())
	// Remove all cookies from request
	req.Header.Set("Cookie", "")
	// Add the required cookies back in the request
	for _, cookieName := range c.keepCookies {
		foundCookie, err := originalRequest.Cookie(cookieName)
		if err == http.ErrNoCookie {
			continue
		}
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		req.AddCookie(foundCookie)
	}
	c.next.ServeHTTP(rw, req)
}
