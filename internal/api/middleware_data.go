package api

import (
	"github.com/Krab1o/avito-backend-assignment-2025/internal/shared"
	errs "github.com/Krab1o/avito-backend-assignment-2025/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

func GetTokenID(c *gin.Context) (int64, error) {
	idRaw, exists := c.Get(shared.UserIDJsonName)
	if !exists {
		return 0, errs.NewServiceError(ErrorInternalServerError, nil)
	}
	idInt := idRaw.(int64)
	return idInt, nil 
}