package user

import (
	"net/http"

	"github.com/Krab1o/avito-backend-assignment-2025/internal/api"
	"github.com/Krab1o/avito-backend-assignment-2025/internal/api/user/converter"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Info(c *gin.Context) {
	id, err := api.GetTokenID(c)
	if err != nil {
		api.HandleError(c, err)
	}
	ctx := c.Request.Context()
	infoService, err := h.infoService.Info(ctx, id)
	if err != nil {
		api.HandleError(c, err)
		return
	}
	infoDTO := converter.InfoServiceToDTO(infoService)
	c.JSON(http.StatusOK, infoDTO)
	return
}