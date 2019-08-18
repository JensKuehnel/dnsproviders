// Package joker adapts the lego Joker DNS
// provider for Caddy. Importing this package plugs it in.
package joker

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/providers/dns/joker"
)

func init() {
	caddytls.RegisterDNSProvider("joker", NewDNSProvider)
}

// NewDNSProvider returns a new Joker DNS challenge provider.
// The credentials are detected automatically; see underlying
// package docs for details:
// https://godoc.org/github.com/go-acme/lego/providers/dns/joker
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return joker.NewDNSProvider()
	default:
		return nil, errors.New("invalid credentials length")
	}
}
