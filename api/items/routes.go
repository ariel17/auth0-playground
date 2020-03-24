package items

import (
	"github.com/ariel17/auth0-playground/api/auth"
	"github.com/gin-gonic/gin"
)

// AddRoutes appends into the indicated engine the items' package routes.
func AddRoutes(r *gin.Engine) {
	u := r.Group("/items")
	u.POST("/", auth.ValidateToken(), createItemController)
	u.GET("/", auth.ValidateToken(), getAllItemsController)
	u.GET("/:id", auth.ValidateToken(), getItemController)
	u.DELETE("/:id", auth.ValidateToken(), deleteItemController)
}
