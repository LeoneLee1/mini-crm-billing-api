package exampletocopy

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Impl(c *gin.Context) {

	httpresputils.HttpRespOK(c, nil, nil, nil)
}
