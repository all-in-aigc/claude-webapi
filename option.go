package claude

import (
	"time"
)

// Options for request client
type Options struct {
	// Debug is used to output debug message
	Debug bool
	// Timeout is used to end http request after timeout duration
	Timeout time.Duration
	// Proxy is used to proxy request
	Proxy string
	// SessionKey is used to set authorization key
	SessionKey string
	// Model is the chat model
	Model string
	// BaseUri is the api base uri
	BaseUri string
	// UserAgent is used to set user agent in header
	UserAgent string
	// Orgid is user's uuid
	Orgid string
}

// Option is used to set custom option
type Option func(*Client)

// WithDebug is used to output debug message
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.opts.Debug = debug
	}
}

// WithTimeout is used to set request timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.opts.Timeout = timeout
	}
}

// WithProxy is used to set request proxy
func WithProxy(proxy string) Option {
	return func(c *Client) {
		c.opts.Proxy = proxy
	}
}

// WithSessionKey is used to set session key in cookie
func WithSessionKey(sessionKey string) Option {
	return func(c *Client) {
		c.opts.SessionKey = sessionKey
	}
}

// WithModel is used to set chat model
func WithModel(model string) Option {
	return func(c *Client) {
		c.opts.Model = model
	}
}

// WithBaseUri is used to set api base uri
func WithBaseUri(baseUri string) Option {
	return func(c *Client) {
		c.opts.BaseUri = baseUri
	}
}

// WithUserAgent is used to set user_agent
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.opts.UserAgent = userAgent
	}
}

// WithOrgid is used to set orgid
func WithOrgid(orgid string) Option {
	return func(c *Client) {
		c.opts.Orgid = orgid
	}
}
