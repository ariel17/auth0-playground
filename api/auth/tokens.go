package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ariel17/auth0-playground/api/config"
	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

// Claims represents the metadata contained in token.
type Claims struct {
	ID            string
	Nickname      string
	GivenName     string
	FamilyName    string
	Picture       string
	Groups        []string
	Roles         []string
	Permissions   []string
	Email         string
	EmailVerified bool
}

var (
	validator      *auth0.JWTValidator
	groupsKey      = config.ApplicationURL + "/roles"
	rolesKey       = config.ApplicationURL + "/roles"
	permissionsKey = config.ApplicationURL + "/permissions"
)

// ValidateToken TODO
// See: https://github.com/auth0-community/auth0-go#example
func ValidateToken() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		token, err := validator.ValidateRequest(c.Request)
		if err != nil {
			fmt.Println("Token is not valid:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			c.Abort()
			return
		}

		rawClaims := map[string]interface{}{}
		err = validator.Claims(c.Request, token, &rawClaims)
		if err != nil {
			fmt.Println("Invalid claims:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		claims := Claims{
			ID:            rawClaims["sub"].(string),
			Nickname:      rawClaims["nickname"].(string),
			GivenName:     rawClaims["given_name"].(string),
			FamilyName:    rawClaims["family_name"].(string),
			Picture:       rawClaims["picture"].(string),
			Groups:        parseClaimArray(rawClaims[groupsKey].([]interface{})),
			Roles:         parseClaimArray(rawClaims[rolesKey].([]interface{})),
			Permissions:   parseClaimArray(rawClaims[permissionsKey].([]interface{})),
			Email:         rawClaims["email"].(string),
			EmailVerified: rawClaims["email_verified"].(bool),
		}

		c.Set("claims", &claims)
		c.Next()
	})
}

// GetClaims extract token claims from context.
func GetClaims(c *gin.Context) (*Claims, error) {
	v, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("claims not found")
	}
	return v.(*Claims), nil
}

func parseClaimArray(values []interface{}) []string {
	newValues := []string{}
	for _, v := range values {
		newValues = append(newValues, v.(string))
	}
	return newValues
}

func newValidator(tenantDomain, audience string) *auth0.JWTValidator {
	domain := "https://" + tenantDomain + "/"
	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: domain + ".well-known/jwks.json"}, nil)
	configuration := auth0.NewConfiguration(client, []string{audience}, domain, jose.RS256)
	return auth0.NewValidator(configuration, nil)
}

func init() {
	validator = newValidator(config.Auth0Domain, config.Auth0Audience)
}
