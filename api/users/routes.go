package users

import (
	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/gin-gonic/gin"
)

// AddRoutes appends into the indicated engine the users' package routes.
func AddRoutes(r *gin.Engine) {
	u := r.Group("/users")
	u.POST("/", auth.ValidateToken(), createUser)
}
