package profileupdate

import (
	"errors"
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/features/public/auth/profile_update/body"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	userID, ok := jwtutils.GetCurrentUserID(c)
	if !ok {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	var req body.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		httpresputils.HttpRespBadRequest(c, &msg)
		return
	}

	user, err := h.usecase.Update(c.Request.Context(), userID, req)
	if err != nil {
		msg := err.Error()
		if errors.Is(err, ErrUserNotFound) {
			httpresputils.HttpRespNotFound(c, &msg)
		} else {
			httpresputils.HttpRespBadRequest(c, &msg)
		}
		return
	}

	msg := "Profile updated"
	httpresputils.HttpRespOK(c, user, nil, &msg)
}
