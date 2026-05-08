package customerlist

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/customer/customer_list/body"
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
	if limit > 100 {
		limit = 100
	}
	status := c.Query("status")
	if status != "" && status != string(models.CustomerStatusActive) && status != string(models.CustomerStatusInActive) {
		msg := "Invalid status value"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	userID, exists := jwtutils.GetCurrentUserID(c)
	if !exists {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}
	requesterID := userID
	requesterRole, _ := c.Get("role")

	result, err := h.usecase.List(c.Request.Context(), body.ListFilter{
		Search:        c.Query("search"),
		Status:        status,
		Page:          page,
		Limit:         limit,
		RequesterID:   requesterID,
		RequesterRole: requesterRole.(string),
	})
	if err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	httpresputils.HttpRespOK(c, result, nil, nil)
}
