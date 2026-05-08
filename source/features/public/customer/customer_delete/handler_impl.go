package customerdelete

import (
	"errors"
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	customerID := c.Param("id")

	if err := h.usecase.Delete(c.Request.Context(), customerID); err != nil {
		msg := err.Error()
		if errors.Is(err, ErrCustomerNotFound) {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Customer deleted"
	httpresputils.HttpRespOK(c, nil, nil, &msg)
}
