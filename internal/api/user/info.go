package user

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/converter"
	"github.com/gin-gonic/gin"
)

//TODO: add validation
func (h *Handler) Info(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := api.ConversionID(c)
	ctx := c.Request.Context()
	infoService, err := h.infoService.Info(ctx, id)
	if err != nil {
		api.HandleServiceError(c, err)
		return
	}
	infoDTO := converter.InfoServiceToDTO(infoService)
	c.JSON(http.StatusOK, infoDTO)
	return
}