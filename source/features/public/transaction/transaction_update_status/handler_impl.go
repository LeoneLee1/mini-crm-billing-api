package transactionupdatestatus

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type updateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func (h *Handler) Impl(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		msg := "invalid transaction id"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	var req updateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := "invalid request body: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	if err := h.usecase.updateStatus(c.Request.Context(), id, req.Status); err != nil {
		msg := err.Error()
		if msg == "transaction not found" {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Transaction status updated successfully"
	httpresputils.HttpRespOK(c, nil, nil, &msg)
}
