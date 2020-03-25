package users

import (
	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/gin-gonic/gin"
)

// AddRoutes appends into the indicated engine the users' package routes.
func AddRoutes(r *gin.Engine) {
	u := r.Group("/users")
	u.POST("/", auth.ValidateToken([]string{"write:users"}), createUserController)
	u.GET("/", auth.ValidateToken([]string{"list:users"}), getAllUsersController)
	u.GET("/:id", auth.ValidateToken([]string{"read:users"}), getUserController)
	u.DELETE("/:id", auth.ValidateToken([]string{"write:users"}), deleteUserController)
}
