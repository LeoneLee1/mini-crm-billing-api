package productcreate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Unit        string  `json:"unit" binding:"required"`
	Category    string  `json:"category"`
	IsActive    *bool   `json:"is_active"`
}

func (h *Handler) Impl(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	userIDStr, exists := c.Get("id")
	if !exists {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	createdBy, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		msg := "Invalid user ID"
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	product, err := h.usecase.Create(c.Request.Context(), createdBy, CreateRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Unit:        req.Unit,
		Category:    req.Category,
		IsActive:    req.IsActive,
	})
	if err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	msg := "Product created"
	httpresputils.HttpRespCreated(c, product, &msg)
}
