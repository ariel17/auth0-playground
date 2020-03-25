package items

import (
	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/gin-gonic/gin"
)

// AddRoutes appends into the indicated engine the items' package routes.
func AddRoutes(r *gin.Engine) {
	u := r.Group("/items")
	u.POST("/", auth.ValidateToken([]string{"write:items"}), createItemController)
	u.GET("/", auth.ValidateToken([]string{"list:items"}), getAllItemsController)
	u.GET("/:id", auth.ValidateToken([]string{"read:items"}), getItemController)
	u.DELETE("/:id", auth.ValidateToken([]string{"write:items"}), deleteItemController)
}
