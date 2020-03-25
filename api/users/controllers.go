package users

import (
	"net/http"
	"time"

	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/gin-gonic/gin"
)

var (
	usersStorage = map[string]*user{}
)

func createUserController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, exists := usersStorage[claims.ID]
	if exists {
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
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	onlyUsers := []*user{}
	if claims.IsAdmin {
		for _, v := range usersStorage {
			onlyUsers = append(onlyUsers, v)
		}
	}
	c.JSON(http.StatusOK, onlyUsers)
}

func getUserController(c *gin.Context) {
	claims, err := auth.GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	user, exists := usersStorage[id]
	if exists && (user.ID == claims.ID || claims.IsAdmin) {
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
	id := c.Param("id")
	user, exists := usersStorage[id]
	if exists && (user.ID == claims.ID || claims.IsAdmin) {
		delete(usersStorage, id)
		now := time.Now()
		user.DeletedAt = &now
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

func saveNewUser(claims *auth.Claims) (*user, error) {
	user := user{
		ID:          claims.ID,
		Nickname:    claims.Nickname,
		GivenName:   claims.GivenName,
		FamilyName:  claims.FamilyName,
		AvatarURL:   claims.Picture,
		Groups:      claims.Groups,
		Roles:       claims.Roles,
		Permissions: claims.Permissions,
		CreatedAt:   time.Now(),
		Enabled:     true,
	}
	user.Email.Address = claims.Email
	user.Email.IsVerified = claims.EmailVerified
	usersStorage[claims.ID] = &user
	return &user, nil
}
