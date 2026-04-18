package productupdate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Unit        *string  `json:"unit"`
	Category    *string  `json:"category"`
	IsActive    *bool    `json:"is_active"`
}

func (h *Handler) Impl(c *gin.Context) {
	id := c.Param("id")

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	product, err := h.usecase.Update(c.Request.Context(), id, UpdateRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Unit:        req.Unit,
		Category:    req.Category,
		IsActive:    req.IsActive,
	})
	if err != nil {
		msg := err.Error()
		if msg == "product not found" {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Product updated"
	httpresputils.HttpRespOK(c, product, nil, &msg)
}
