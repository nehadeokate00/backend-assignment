package handlers

import (
    "backend-assignment/database"
    "backend-assignment/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

    email := c.MustGet("email").(string)

    var user models.User

    database.DB.Where("email = ?", email).First(&user)

    c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {

    role := c.MustGet("role").(string)

    if role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{
            "error": "Access denied",
        })
        return
    }

    var users []models.User

    database.DB.Find(&users)

    c.JSON(http.StatusOK, users)
}