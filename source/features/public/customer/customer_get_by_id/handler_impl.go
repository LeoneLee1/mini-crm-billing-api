package customergetbyid

import (
	"errors"
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

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

	customer, err := h.usecase.GetByID(c.Request.Context(), customerID)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, ErrCustomerNotFound) {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	httpresputils.HttpRespOK(c, customer, nil, nil)
}
