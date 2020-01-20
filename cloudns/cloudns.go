// Package cloudns adapts the lego CloudNS
// provider for Caddy. Importing this package plugs it in.
package cloudns

import (
	"errors"

	"github.com/caddyserver/caddy/caddytls"
	"github.com/go-acme/lego/v3/providers/dns/cloudns"
)

func init() {
	caddytls.RegisterDNSProvider("cloudns", NewDNSProvider)
}

// NewDNSProvider returns a new CloudNS DNS challenge provider.
// The credentials are interpreted as follows:
//
// len(0): use credentials from environment
// len(2): credentials[0] = Api user ID
//         credentials[1] = API password
func NewDNSProvider(credentials ...string) (caddytls.ChallengeProvider, error) {
	switch len(credentials) {
	case 0:
		return cloudns.NewDNSProvider()
	case 2:
		config := cloudns.NewDefaultConfig()
		config.AuthID = credentials[0]
		config.AuthPassword = credentials[1]
		return cloudns.NewDNSProviderConfig(config)
	default:
		return nil, errors.New("invalid credentials length")
	}
}
