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
	c.Header("Content-Type", "application/json")
	buyerID := api.ConversionID(c)
	ctx := c.Request.Context()
	newBuying := &dto.Buying{
		Name: c.Param(api.ParamBuying),
	}
	if newBuying.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorNoParamSpecified})
	}

	//TODO: everything before is validation
	model := converter.BuyingDTOToService(newBuying, buyerID)
	err := h.buyingService.Buy(ctx, model)
	log.Println("hehe", err)
	if err != nil {
		api.HandleServiceError(c, err)
		// c.JSON(http.StatusInternalServerError, gin.H{api.ErrorField: api.ErrorInternalServerError})
		return
	}
	return
}