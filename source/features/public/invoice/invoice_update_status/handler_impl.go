package invoiceupdatestatus

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
		errMSG := "invalid invoice id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	var req updateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMSG := "invalid request body: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	invoice, err := h.usecase.updateStatus(c.Request.Context(), id, req.Status)
	if err != nil {
		errMSG := err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	msg := "Invoice status updated successfully"
	httpresputils.HttpRespOK(c, invoice, nil, &msg)
}
