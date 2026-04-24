package profileupdate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	id, ok := jwtutils.GetCurrentUserID(c)
	if !ok {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	user, err := h.usecase.update(c.Request.Context(), id, updateRequest{
		Name:  req.Name,
		Email: req.Email,
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

	msg := "Profile updated"
	httpresputils.HttpRespOK(c, user, nil, &msg)
}
