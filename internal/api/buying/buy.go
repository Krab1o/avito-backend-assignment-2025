package buying

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO: Think of generalization of itemName (maybe add struct)
func (h *Handler) Buy(c *gin.Context) {
	ctx := c.Request.Context()
	item := c.Param("item")
	if item == "" {
		log.Printf("Handler buying: no param specified")
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorNoParamSpecified})
	}
	err := h.service.Buy(ctx, item)
	if err != nil {
		log.Printf("Handler buying: %v", err)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{"errors": errorInternalServerError})
		return
	}
	c.Status(http.StatusOK)
	return
}