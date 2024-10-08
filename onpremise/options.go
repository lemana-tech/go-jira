package onpremise

import (
	"log/slog"
	"net/http"
)

// Opt func type
type Opt func(c *Client)

// WithHTTPClient sets the http.Client to use for API requests.
func WithHTTPClient(hc *http.Client) Opt {
	return func(c *Client) {
		c.client = hc
	}
}

// WithLogger sets the slog instance.
func WithLogger(l *slog.Logger) Opt {
	return func(c *Client) {
		c.logger = l
	}
}
