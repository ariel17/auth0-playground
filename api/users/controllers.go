package users

import (
	"net/http"
	"time"

	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/ariel17/auth0-playground/api/config"
	"github.com/gin-gonic/gin"
)

var (
	groupsKey      = config.ApplicationURL + "/roles"
	rolesKey       = config.ApplicationURL + "/roles"
	permissionsKey = config.ApplicationURL + "/permissions"
	usersStorage   = map[string]*user{}
)

func createUserController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if _, exists := getUser(claims); exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
		return
	}
	user, err := saveNewUser(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func getAllUsersController(c *gin.Context) {
	onlyUsers := []*user{}
	for _, v := range usersStorage {
		onlyUsers = append(onlyUsers, v)
	}
	c.JSON(http.StatusOK, onlyUsers)
}

func getUserController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user, exists := getUser(claims); exists {
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func deleteUserController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, exists := getUser(claims)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	delete(usersStorage, user.ID)
	now := time.Now()
	user.DeletedAt = &now
	c.JSON(http.StatusOK, user)
}

func getUser(claims auth.Claims) (*user, bool) {
	id := claims["sub"].(string)
	user, exists := usersStorage[id]
	return user, exists
}

func saveNewUser(claims auth.Claims) (*user, error) {
	id := claims["sub"].(string)

	user := user{
		ID:          id,
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
	usersStorage[id] = &user
	return &user, nil
}

func parseClaimArray(values []interface{}) []string {
	newValues := []string{}
	for _, v := range values {
		newValues = append(newValues, v.(string))
	}
	return newValues
}
