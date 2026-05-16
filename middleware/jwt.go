package middleware

import (
    "backend-assignment/utils"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        authHeader := c.GetHeader("Authorization")

        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "No token found",
            })
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return utils.SecretKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid token",
            })
            c.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)

        c.Set("email", claims["email"])
        c.Set("role", claims["role"])

        c.Next()
    }
}