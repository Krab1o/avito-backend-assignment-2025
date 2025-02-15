package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	shared "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func (c *gin.Context) {
		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.JSON(http.StatusUnauthorized, gin.H{api.FieldError: api.ErrorUnauthorized})
				c.Abort()
				return
			}
	
		// Разбираем заголовок (должен быть вида "Bearer <token>")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{api.FieldError: api.ErrorInvalidAuthorizationHeader})
			c.Abort()
			return
		}
	
		tokenString := parts[1]
	
		// Парсим и проверяем JWT
		claims := &shared.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method")
				}
				return jwtSecret, nil
			})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{api.FieldError: api.ErrorToken})
			c.Abort()
			return
		}

		// Передаём user_id в контекст Gin
		c.Set(shared.UserIDJsonName, claims.UserID)
		c.Next()
	}
}