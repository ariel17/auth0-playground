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
type Claims map[string]interface{}

var validator *auth0.JWTValidator

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

		claims := Claims{}
		err = validator.Claims(c.Request, token, &claims)
		if err != nil {
			fmt.Println("Invalid claims:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	})
}

// GetClaims extract token claims from context.
func GetClaims(c *gin.Context) (Claims, error) {
	v, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("claims not found")
	}
	return v.(Claims), nil
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
