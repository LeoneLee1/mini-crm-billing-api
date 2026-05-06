package invoicedelete

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Impl(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		errMSG := "invalid invoice id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	if err := h.usecase.delete(c.Request.Context(), id); err != nil {
		errMSG := err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	msg := "Invoice deleted successfully"
	httpresputils.HttpRespOK(c, nil, nil, &msg)
}
