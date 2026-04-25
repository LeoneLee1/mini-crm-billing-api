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
	// redis *redis.Client
}

type healthResponse struct {
	Status     string    `json:"status"`
	AppName    string    `json:"app_name"`
	Time       time.Time `json:"timestamp"`
	AppVersion string    `json:"app_version"`
	DB         string    `json:"db"`
	// Redis      string    `json:"redis"`
}

// func CheckHealth(db *gorm.DB, redisClient *redis.Client) *Handler {
func CheckHealth(db *gorm.DB) *Handler {
	// return &Handler{db: db, redis: redisClient}
	return &Handler{db: db}
}

var (
	cfg = config.Load()
)

func (h *Handler) Check(c *gin.Context) {
	dbStatus := "up"
	// redisStatus := "up"
	overall := "OK"

	sqlDB, err := h.db.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "down"
		overall = "Unhealthy"
	}

	// if h.redis == nil || h.redis.Ping(context.Background()).Err() != nil {
	// 	redisStatus = "down"
	// 	overall = "Unhealthy"
	// }

	statusCode := http.StatusOK
	if overall != "OK" {
		statusCode = http.StatusServiceUnavailable
	}

	c.JSON(statusCode, healthResponse{
		Status: overall,
		DB:     dbStatus,
		// Redis:      redisStatus,
		AppName:    cfg.AppName,
		Time:       time.Now(),
		AppVersion: cfg.AppVersion,
	})
}
