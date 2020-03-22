package config

import "os"

var (
	// Auth0Domain is your TENANT + auth0.com.
	Auth0Domain string

	// Auth0Audience is the Auth0 token IDENTIFIER:
	// If using an SPA: CLIENT ID.
	// If using an API: AUDIENCE (Identifier).
	Auth0Audience string
)

func init() {
	Auth0Domain = os.Getenv("AUTH0_DOMAIN")
	Auth0Audience = os.Getenv("AUTH0_AUDIENCE")
}
