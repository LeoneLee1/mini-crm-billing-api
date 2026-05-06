package invoicegetbyid

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

	invoice, err := h.usecase.getByID(c.Request.Context(), id)
	if err != nil {
		errMSG := err.Error()
		httpresputils.HttpRespNotFound(c, &errMSG)
		return
	}

	httpresputils.HttpRespOK(c, invoice, nil, nil)
}
