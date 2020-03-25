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

var (
	validator      *auth0.JWTValidator
	groupsKey      = config.ApplicationURL + "/roles"
	rolesKey       = config.ApplicationURL + "/roles"
	permissionsKey = config.ApplicationURL + "/permissions"
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

// HasPermissions check if the indicated list of permissions exists in claims.
func (c *Claims) HasPermissions(required []string) bool {
	fmt.Println("permissions:", c.Permissions)
	fmt.Println("required:", required)
	for _, r := range required {
		var found bool
		for _, p := range c.Permissions {
			if r == p {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// ValidateToken TODO
// See: https://github.com/auth0-community/auth0-go#example
func ValidateToken(requiredPermissions []string) gin.HandlerFunc {
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

		claims := createClaims(rawClaims)
		if !claims.HasPermissions(requiredPermissions) {
			fmt.Printf("Invalid permissions: current=%v, required=%v", claims.Permissions, requiredPermissions)
			c.JSON(http.StatusForbidden, gin.H{"error": "invalid permissions"})
			c.Abort()
			return
		}
		c.Set("claims", claims)

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

func createClaims(claims map[string]interface{}) *Claims {
	c := Claims{
		ID:            claims["sub"].(string),
		Nickname:      claims["nickname"].(string),
		GivenName:     claims["given_name"].(string),
		FamilyName:    claims["family_name"].(string),
		Picture:       claims["picture"].(string),
		Groups:        parseClaimArray(claims[groupsKey].([]interface{})),
		Roles:         parseClaimArray(claims[rolesKey].([]interface{})),
		Permissions:   parseClaimArray(claims[permissionsKey].([]interface{})),
		Email:         claims["email"].(string),
		EmailVerified: claims["email_verified"].(bool),
	}
	return &c
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
