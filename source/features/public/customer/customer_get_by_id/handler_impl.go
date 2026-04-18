package customergetbyid

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	id := c.Param("id")

	customer, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		msg := err.Error()
		httpresputils.HttpRespNotFound(c, &msg)
		return
	}

	httpresputils.HttpRespOK(c, customer, nil, nil)
}
