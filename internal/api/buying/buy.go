package buying

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/buying/dto"
	"github.com/gin-gonic/gin"
)

func BuyingValidation(c *gin.Context) *dto.Buying {
	buying := &dto.Buying{
		Name: c.Param(api.ParamBuying),
	}
	if buying.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorNoParamSpecified})
		return nil
	}
	return buying
}

func (h *Handler) Buy(c *gin.Context) {
	buyerID, err := api.GetTokenID(c)
	if err != nil {
		api.HandleError(c, err)
		return
	}
	ctx := c.Request.Context()
	buying := BuyingValidation(c) 
	if buying == nil {
		return
	}
	model := converter.BuyingDTOToService(buying, buyerID)
	err = h.buyingService.Buy(ctx, model)
	if err != nil {
		api.HandleError(c, err)
		return
	}
	return
}