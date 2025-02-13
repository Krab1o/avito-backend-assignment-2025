package transaction

import (
	"log"
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/converter"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/transaction/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCoin(c *gin.Context) {	
	ctx := c.Request.Context()
	newTransaction := &dto.Transaction{}
	err := c.ShouldBindJSON(newTransaction)
	if err != nil {
		log.Printf("Handler sendCoin: failed to convert body, %v", err)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{"errors" : api.ErrorBadRequest})
		return
	}
	model := converter.TransactionDTOToService(newTransaction)
	err = h.transactionService.SendCoin(ctx, model)
	if err != nil {
		log.Printf("Handler sendCoin: internal error %v", err)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": api.ErrorInternalServerError})
		return
	}
	c.Status(http.StatusOK)
}