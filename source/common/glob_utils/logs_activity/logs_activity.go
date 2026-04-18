package logsactivity

import (
	"encoding/json"
	"mini-crm-billing-api/source/common/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppendLogsActivity(ctx *gin.Context, db *gorm.DB, before, after any) {

	time.Sleep(1 * time.Second)
	logActivity := ctx.GetString("log_id")

	stmt := db.Model(&models.LogActivityModel{}).Where("id = ?", logActivity)

	beforeJSON, err := json.Marshal(before)
	if before != nil && err == nil {
		stmt.Updates(map[string]any{
			"data_before": string(beforeJSON),
		})

	}

	afterJSON, err := json.Marshal(after)
	if err == nil && after != nil {
		stmt.Updates(map[string]any{
			"data_after": string(afterJSON),
		})
	}
}
