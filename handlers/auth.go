package handlers

import (
    "backend-assignment/database"
    "backend-assignment/models"
    "backend-assignment/utils"
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    user.Password = string(hashedPassword)

    database.DB.Create(&user)

    c.JSON(http.StatusOK, gin.H{
        "message": "User created successfully",
    })
}

func Login(c *gin.Context) {

    var input models.User
    var user models.User

    c.ShouldBindJSON(&input)

    database.DB.Where("email = ?", input.Email).First(&user)

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid credentials",
        })
        return
    }

    token, _ := utils.GenerateJWT(user.Email, user.Role)

    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}