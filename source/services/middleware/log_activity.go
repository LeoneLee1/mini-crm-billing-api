package middleware

import (
	"fmt"
	"mini-crm-billing-api/source/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func LogActivityMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.GetString("id")
		if userID == "" {
			return
		}

		status := c.Writer.Status()
		action := fmt.Sprintf("[%v][%s] %s", http.StatusText(status), c.Request.Method, c.Request.URL.Path)

		userIDParse, err := uuid.Parse(userID)
		if err != nil {
			return
		}

		log := models.LogActivityModel{
			ID:     uuid.New(),
			UserID: userIDParse,
			Action: action,
		}
		c.Set("log_id", log.ID.String())
		db.Create(&log)

		c.Next()
	}
}
