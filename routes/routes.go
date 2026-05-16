package routes

import (
    "backend-assignment/handlers"
    "backend-assignment/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

    router.POST("/signup", handlers.Signup)
    router.POST("/login", handlers.Login)

    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())

    protected.GET("/profile", handlers.Profile)
    protected.GET("/users", handlers.GetUsers)
}