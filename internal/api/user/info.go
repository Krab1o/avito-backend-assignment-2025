package user

import (
	"log"
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/converter"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Info(c *gin.Context) {
	ctx := c.Request.Context()
	info, err := h.infoService.Info(ctx)
	if err != nil {
		log.Printf("Handler info: %v", err)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{api.FieldError: api.ErrorInternalServerError})
		return
	}
	//TODO: everything before is validation
	dtoInfo := converter.InfoServiceToDTO(info)
	c.JSON(http.StatusOK, dtoInfo)
	return
}