package customercreate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Status  string `json:"status"`
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

	customer, err := h.usecase.Create(c.Request.Context(), createdBy, CreateRequest{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		Status:  req.Status,
	})
	if err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	msg := "Customer created"
	httpresputils.HttpRespCreated(c, customer, &msg)
}
