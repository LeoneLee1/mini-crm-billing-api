package middleware

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			errMSG := "Admin only"
			httpresputils.HttpResponseForbidden(c, &errMSG)
			c.Abort()
			return
		}

		if role != "admin" {
			errMSG := "Admin only"
			httpresputils.HttpResponseForbidden(c, &errMSG)
			c.Abort()
			return
		}

		c.Next()
	}
}
