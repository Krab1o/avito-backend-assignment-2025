package params

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/gin-gonic/gin"
)

func NoParamsMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		if c.Request.URL.RawQuery != "" {
			c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorParametersNotAllowed})
			c.Abort()
			return
		}
	}
}