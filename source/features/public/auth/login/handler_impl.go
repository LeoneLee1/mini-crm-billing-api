package login

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Impl(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	result, err := h.usecase.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		msg := err.Error()
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	msg := "Login successful"
	httpresputils.HttpRespOK(c, result, nil, &msg)
}
