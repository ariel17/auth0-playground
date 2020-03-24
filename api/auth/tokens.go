package auth

import (
	"fmt"
	"net/http"

	"github.com/ariel17/auth0-playground/api/config"
	auth0 "github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

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

		claims := map[string]interface{}{}
		err = validator.Claims(c.Request, token, &claims)
		if err != nil {
			fmt.Println("Invalid claims:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		fmt.Println("Claims:", claims)
		c.Next()
	})
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
