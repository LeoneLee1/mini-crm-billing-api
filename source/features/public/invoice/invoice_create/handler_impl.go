package invoicecreate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	"mini-crm-billing-api/source/features/public/invoice/invoice_create/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	var req body.InvoiceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMSG := "invalid request body: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	invoice, err := h.usecase.Create(c.Request.Context(), &req)
	if err != nil {
		errMSG := err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	msg := "Invoice created successfully"
	httpresputils.HttpRespCreated(c, invoice, &msg)
}
