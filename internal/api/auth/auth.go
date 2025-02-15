package auth

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/dto"
	"github.com/gin-gonic/gin"
)

//TODO: validation
func (h *Handler) Auth(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	ctx := c.Request.Context()
	creds := &dto.UserCreds{}
	err := c.ShouldBindJSON(creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorBadRequest})
		return
	}
	// everything before is validation
	serviceCreds := converter.CredsDTOToService(creds)
	token, err := h.authService.Auth(ctx, serviceCreds)
	if err != nil {
		api.HandleServiceError(c, err)
		// c.JSON(http.StatusUnauthorized, gin.H{api.ErrorField: api.ErrorUnauthorized})
		return
	}
	c.JSON(http.StatusOK, gin.H{api.FieldToken: token})
}