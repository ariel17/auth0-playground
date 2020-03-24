package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/ariel17/auth0-playground/api/config"
	"github.com/gin-gonic/gin"
)

var (
	groupsKey      = config.ApplicationURL + "/roles"
	rolesKey       = config.ApplicationURL + "/roles"
	permissionsKey = config.ApplicationURL + "/permissions"
	usersStorage   = []*user{}
)

func createUser(c *gin.Context) {
	user, err := saveNewUser(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func saveNewUser(c *gin.Context) (*user, error) {
	v, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("claims not found")
	}
	claims := v.(map[string]interface{})
	user := user{
		ID:          claims["sub"].(string),
		Nickname:    claims["nickname"].(string),
		GivenName:   claims["given_name"].(string),
		FamilyName:  claims["family_name"].(string),
		AvatarURL:   claims["picture"].(string),
		Groups:      parseClaimArray(claims[groupsKey].([]interface{})),
		Roles:       parseClaimArray(claims[rolesKey].([]interface{})),
		Permissions: parseClaimArray(claims[permissionsKey].([]interface{})),
		CreatedAt:   time.Now(),
		Enabled:     true,
	}
	user.Email.Address = claims["email"].(string)
	user.Email.IsVerified = claims["email_verified"].(bool)

	usersStorage = append(usersStorage, &user)
	return &user, nil
}

func parseClaimArray(values []interface{}) []string {
	newValues := []string{}
	for _, v := range values {
		newValues = append(newValues, v.(string))
	}
	return newValues
}
