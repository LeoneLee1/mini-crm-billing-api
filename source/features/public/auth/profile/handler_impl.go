package profile

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	user, err := h.usecase.GetProfile(c.Request.Context(), userID.(string))
	if err != nil {
		msg := "User not found"
		httpresputils.HttpRespNotFound(c, &msg)
		return
	}

	httpresputils.HttpRespOK(c, user, nil, nil)
}
