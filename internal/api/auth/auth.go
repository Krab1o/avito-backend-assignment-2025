package auth

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(c *gin.Context) {
	ctx := c.Request.Context()
	creds := &dto.UserCreds{}
	err := c.ShouldBindJSON(creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": api.ErrorBadRequest})
		return
	}
	serviceCreds := converter.CredsDTOToService(creds)
	h.authService.Auth(ctx, serviceCreds)
}