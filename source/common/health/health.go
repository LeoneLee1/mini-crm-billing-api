package health

import (
	"mini-crm-billing-api/source/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

type healthResponse struct {
	Status     string    `json:"status"`
	AppName    string    `json:"app_name"`
	Time       time.Time `json:"timestamp"`
	AppVersion string    `json:"app_version"`
	DB         string    `json:"db"`
}

func CheckHealth(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

var (
	cfg = config.Load()
)

func (h *Handler) Check(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, healthResponse{
			Status:     "Unhealthy",
			DB:         "not connected",
			AppName:    cfg.AppName,
			Time:       time.Now(),
			AppVersion: cfg.AppVersion,
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, healthResponse{
			Status:     "Unhealthy",
			DB:         "dow",
			AppName:    cfg.AppName,
			Time:       time.Now(),
			AppVersion: cfg.AppVersion,
		})
		return
	}

	c.JSON(http.StatusOK, healthResponse{
		Status:     "OK",
		DB:         "up",
		AppName:    cfg.AppName,
		Time:       time.Now(),
		AppVersion: cfg.AppVersion,
	})
}
