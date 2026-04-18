package transactionupdate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	"mini-crm-billing-api/source/features/public/transaction/transaction_update/body"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Impl(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		msg := "invalid transaction id"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	var req body.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := "invalid request body: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	transaction, err := h.usecase.update(c.Request.Context(), id, &req)
	if err != nil {
		msg := err.Error()
		if msg == "transaction not found" {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Transaction updated successfully"
	httpresputils.HttpRespOK(c, transaction, nil, &msg)
}
