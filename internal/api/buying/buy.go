package buying

import (
	"log"
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Buy(c *gin.Context) {
	ctx := c.Request.Context()
	dto := &dto.Buying{
		Name: c.Param("item"),
	}
	if dto.Name == "" {
		log.Printf("Handler buying: no param specified")
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{"errors": api.ErrorNoParamSpecified})
	}
	model := converter.BuyingDTOToService(dto)
	err := h.buyingService.Buy(ctx, model)
	if err != nil {
		log.Printf("Handler buying: %v", err)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": api.ErrorInternalServerError})
		return
	}
	c.Status(http.StatusOK)
	return
}