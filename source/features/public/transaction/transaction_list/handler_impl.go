package transactionlist

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	customerID := c.Query("customer_id")
	status := c.Query("status")

	result, err := h.usecase.list(c.Request.Context(), customerID, status, page, limit)
	if err != nil {
		errMSG := err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, result, nil, nil)
}
