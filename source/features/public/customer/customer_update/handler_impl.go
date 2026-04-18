package customerupdate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Name    *string `json:"name"`
	Email   *string `json:"email"`
	Phone   *string `json:"phone"`
	Address *string `json:"address"`
	Status  *string `json:"status"`
}

func (h *Handler) Impl(c *gin.Context) {
	id := c.Param("id")

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	customer, err := h.usecase.Update(c.Request.Context(), id, UpdateRequest{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Status:  req.Status,
	})
	if err != nil {
		msg := err.Error()
		if msg == "customer not found" {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Customer updated"
	httpresputils.HttpRespOK(c, customer, nil, &msg)
}
