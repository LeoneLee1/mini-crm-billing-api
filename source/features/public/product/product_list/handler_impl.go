package productlist

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

	var isActive *bool
	if v := c.Query("is_active"); v != "" {
		b, err := strconv.ParseBool(v)
		if err == nil {
			isActive = &b
		}
	}

	result, err := h.usecase.List(c.Request.Context(), ListFilter{
		Search:   c.Query("search"),
		Category: c.Query("category"),
		IsActive: isActive,
		Page:     page,
		Limit:    limit,
	})
	if err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	httpresputils.HttpRespOK(c, result, nil, nil)
}
