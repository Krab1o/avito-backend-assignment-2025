package api

import (
	shared "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/jwt"
	"github.com/gin-gonic/gin"
)

func ConversionID(c *gin.Context) int64 {
	idRaw, _ := c.Get(shared.UserIDJsonName)
	idInt := idRaw.(int64)
	return idInt 
}