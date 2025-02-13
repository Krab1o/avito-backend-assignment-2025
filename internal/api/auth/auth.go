package auth

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(c *gin.Context) {
	// ctx := c.Request.Context()
	// user := &dto.UserCreds{}
	// err := c.ShouldBindJSON(user)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"errors": api.ErrorBadRequest})
	// 	return
	// }
	// h.authService.Auth(ctx)
}