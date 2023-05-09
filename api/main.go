package main

import (
	"github.com/gin-gonic/gin"
	"./auth"
)

func main() {
    router := gin.Default()

    router.POST("/login", auth.LoginHandler)
    authorized := router.Group("/", atuh.AuthMiddleware())
    {
        authorized.GET("/protected", protectedHandler)
    }

    router.Run(":8080")
}
