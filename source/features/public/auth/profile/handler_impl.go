package profile

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	userrepo "mini-crm-billing-api/source/common/repository/user_repo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {
	userID, exists := jwtutils.GetCurrentUserID(c)
	if !exists {
		msg := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &msg)
		return
	}

	user, err := userrepo.FindByID(c.Request.Context(), h.db, userID)
	if err != nil {
		msg := "User not found"
		httpresputils.HttpRespNotFound(c, &msg)
		return
	}

	httpresputils.HttpRespOK(c, user, nil, nil)
}
