package auth

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/auth/dto"
	"github.com/gin-gonic/gin"
)

func CredsValidation(c *gin.Context) *dto.UserCreds {
	creds := &dto.UserCreds{}
	err := c.ShouldBindJSON(creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorBadRequest})
		return nil
	}
	if creds.Password == "" || creds.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorBadRequest})
		return nil
	}
	return creds
}

func (h *Handler) Auth(c *gin.Context) {
	ctx := c.Request.Context()
	creds := CredsValidation(c)
	if creds == nil {
		return
	}
	serviceCreds := converter.CredsDTOToService(creds)
	token, err := h.authService.Auth(ctx, serviceCreds)
	if err != nil {
		api.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{api.FieldToken: token})
}