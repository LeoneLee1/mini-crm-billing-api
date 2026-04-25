package userupdate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Role  *string `json:"role" binding:"omitempty,oneof=admin staff"`
}

func (h *Handler) Impl(c *gin.Context) {
	id := c.Param("id")

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	user, err := h.usecase.Update(c.Request.Context(), id, UpdateRequest{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	})
	if err != nil {
		msg := err.Error()
		if msg == "user not found" {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "User updated"
	httpresputils.HttpRespOK(c, user, nil, &msg)
}
