package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	domain := "singularmentor-dev.auth0.com"
	audience := "https://brian.sinexiar.com.ar"
	validator = newValidator(domain, audience)

	testCases := []struct {
		name    string
		headers map[string]string
		isValid bool
	}{
		{"missing token", nil, false},
		{"invalid token", map[string]string{"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}, false},
		{"valid token", map[string]string{"authorization": "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlJUVTJSamMwUmtOQ1JFWkNRelkyTXpZMU5rTkdRMFEwTURORk9UYzVNekk1T1VRMVFVTkROUSJ9.eyJpc3MiOiJodHRwczovL3Npbmd1bGFybWVudG9yLWRldi5hdXRoMC5jb20vIiwic3ViIjoiVElaUWFQM0g1dmV6dFFqemFKQ1Njb1MxVjVBWkw5OFZAY2xpZW50cyIsImF1ZCI6Imh0dHBzOi8vYnJpYW4uc2luZXhpYXIuY29tLmFyIiwiaWF0IjoxNTc2NTU5MjkzLCJleHAiOjE1NzY2NDU2OTMsImF6cCI6IlRJWlFhUDNINXZlenRRanphSkNTY29TMVY1QVpMOThWIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIn0.eZmQXppJb06sNNrR9pYRPhp_SHZWzvKmar31wT7q42n_DKq0pAUE94nbNTP5hIo64hoyStwDLf7HcPjV7ou2GManaVVNbNfMVJa4Acc7qICvyCRO-HOPFa5TxK_8PQfSW040SAMMfj0GMvW89PNVrT7JrR_rBrlhT2G6cs3GraZRPCV1zKxecqC5CzfJokqUEJTgCJV0vpFX8F6RAk30iZrvf4WwQ-ziwUDN1ZFkUA62-N0mYHskXEmau5XJPYNFGPrc2O7RlCJSJM03qEjRKliokx_qhwylQr-kgGaHJzcRIg7W6FGKYIaEzv4d6b-K4ayXH2MYCBY4bV82fs2zJA"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, r := gin.CreateTestContext(response)
			r.Use(ValidateToken())
			r.GET("/test", func(c *gin.Context) {
				c.Status(http.StatusOK)
			})
			c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
			if tc.headers != nil {
				for k, v := range tc.headers {
					c.Request.Header.Set(k, v)
				}
			}
			r.ServeHTTP(response, c.Request)
			if tc.isValid {
				assert.Equal(t, http.StatusOK, response.Code)
			} else {
				assert.Equal(t, http.StatusUnauthorized, response.Code)
			}
		})
	}
}
