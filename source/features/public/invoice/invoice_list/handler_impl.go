package invoicelist

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/common/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	status := c.Query("status")
	customerID := c.Query("customer_id")

	// staff hanya bisa lihat invoice dari transaksi yang dia buat
	var createdByFilter *uuid.UUID
	role, _ := c.Get("role")
	if role == string(models.UserRoleStaff) {
		userID, ok := jwtutils.GetCurrentUserID(c)
		if !ok {
			errMSG := "unauthorized"
			httpresputils.HttpResponseUnAuth(c, &errMSG)
			return
		}
		parsed, err := uuid.Parse(userID)
		if err != nil {
			errMSG := "invalid user id"
			httpresputils.HttpRespBadRequest(c, &errMSG)
			return
		}
		createdByFilter = &parsed
	}

	result, err := h.usecase.list(c.Request.Context(), status, customerID, createdByFilter, page, limit)
	if err != nil {
		errMSG := err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, result, nil, nil)
}
