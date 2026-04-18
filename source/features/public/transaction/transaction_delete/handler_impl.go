package transactiondelete

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Impl(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		msg := "invalid transaction id"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	if err := h.usecase.delete(c.Request.Context(), id); err != nil {
		msg := err.Error()
		if msg == "transaction not found" {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Transaction deleted successfully"
	httpresputils.HttpRespOK(c, nil, nil, &msg)
}
