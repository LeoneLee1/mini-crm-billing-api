package main

import (
	"fmt"
	"mini-crm-billing-api/source/common/health"
	"mini-crm-billing-api/source/config"
	"mini-crm-billing-api/source/pkg/db"
	"mini-crm-billing-api/source/pkg/logger"
	pkgredis "mini-crm-billing-api/source/pkg/redis"
	"mini-crm-billing-api/source/services"
	"mini-crm-billing-api/source/services/middleware"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init(true)
	logger.Info().Msg("App Started")

	cfg := config.Load()
	dbConn, err := db.Database(cfg)
	if err != nil {
		logger.Error().Msg("Database connection failed")
		return
	}
	logger.Info().Msg("Database connected")

	redisClient, err := pkgredis.NewClient(cfg)
	if err != nil {
		logger.Error().Msg("Redis connection failed")
		return
	}
	logger.Info().Msg("Redis connected")

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           86400,
	}))
	r.Use(middleware.RequestIDMiddleware())
	r.Use(logger.GinZLogger())
	r.Use(gin.Recovery())

	// HealthCheck /health
	healthHandler := health.CheckHealth(dbConn, redisClient)
	r.GET("/health", healthHandler.Check)

	// Mounting routers
	route_api_v1 := r.Group("api/v1")
	mounthRote := services.NewRouters(dbConn)
	mounthRote.MountRouters(route_api_v1)

	svc := &http.Server{
		Addr:           fmt.Sprintf(":%v", cfg.AppPort),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Run Server
	logger.Info().
		Str("port", fmt.Sprintf("%v", cfg.AppPort)).
		Msg("Server running")

	if err := svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Log.Fatal().Err(err).Msg("Server failed to start")
	}
}
