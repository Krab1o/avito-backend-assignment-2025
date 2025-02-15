package transaction

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/dto"
	"github.com/gin-gonic/gin"
)

//TODO: add validation middleware
func (h *Handler) SendCoin(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	senderID := api.ConversionID(c)
	ctx := c.Request.Context()
	newTransaction := &dto.Transaction{}
	err := c.ShouldBindJSON(newTransaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorBadRequest})
		return
	}
	//TODO: everything before is validation
	model := converter.TransactionDTOToService(newTransaction, senderID)
	err = h.transactionService.SendCoin(ctx, model)
	if err != nil {
		api.HandleServiceError(c, err)
		// c.JSON(http.StatusInternalServerError, gin.H{api.ErrorField: api.ErrorInternalServerError})
		return
	}
	c.Status(http.StatusOK)
}