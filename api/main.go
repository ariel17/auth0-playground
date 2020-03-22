package main

import (
	"github.com/ariel17/auth0-playground/api/users"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	users.AddRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
