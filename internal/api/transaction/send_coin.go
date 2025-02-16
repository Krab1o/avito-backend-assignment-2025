package transaction

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/dto"
	"github.com/gin-gonic/gin"
)

func TransactionValidation(c *gin.Context) *dto.Transaction {
	tx := &dto.Transaction{}
	err := c.ShouldBindJSON(tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorBadRequest})
		return nil
	}
	if tx.Amount <= 0 || tx.ToUser == "" {
		c.JSON(http.StatusBadRequest, gin.H{api.FieldError: api.ErrorBadRequest})
		return nil
	}
	return tx
}

func (h *Handler) SendCoin(c *gin.Context) {
	ctx := c.Request.Context()
	senderID, err := api.GetTokenID(c)
	if err != nil {
		api.HandleError(c, err)
	}
	tx := TransactionValidation(c)
	if tx == nil {
		return
	}
	model := converter.TransactionDTOToService(tx, senderID)
	err = h.transactionService.SendCoin(ctx, model)
	if err != nil {
		api.HandleError(c, err)
		return
	}
	c.Status(http.StatusOK)
}