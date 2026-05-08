package customerupdate

import (
	"errors"
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/features/public/customer/customer_update/body"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Impl(c *gin.Context) {
	customerID := c.Param("id")
	if _, err := uuid.Parse(customerID); err != nil {
		msg := "Invalid customer ID"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	userInfo, ok := jwtutils.GetCurrentUser(c)
	if !ok {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	var req body.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	customer, err := h.usecase.Update(c.Request.Context(), customerID, req, userInfo.ID, userInfo.Role)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, ErrCustomerNotFound) {
			httpresputils.HttpRespNotFound(c, &msg)
		} else if errors.Is(err, ErrForbidden) {
			httpresputils.HttpResponseForbidden(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Customer updated"
	httpresputils.HttpRespOK(c, customer, nil, &msg)
}
