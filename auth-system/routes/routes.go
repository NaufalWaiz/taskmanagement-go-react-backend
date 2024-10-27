package routes

import (
    "project/handlers"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/register", handlers.Register)
    router.POST("/login", handlers.Login)
}