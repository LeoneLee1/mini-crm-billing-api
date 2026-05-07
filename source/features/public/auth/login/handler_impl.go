package login

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	"mini-crm-billing-api/source/features/public/auth/login/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	var req body.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	result, err := h.usecase.Login(c.Request.Context(), h.db, req.Email, req.Password)
	if err != nil {
		msg := err.Error()
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	msg := "Login successful"
	httpresputils.HttpRespOK(c, result, nil, &msg)
}
