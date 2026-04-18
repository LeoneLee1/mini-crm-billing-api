package services

import (
	"mini-crm-billing-api/source/features/public/auth/login"
	"mini-crm-billing-api/source/features/public/auth/profile"
	"mini-crm-billing-api/source/features/public/auth/register"
	"mini-crm-billing-api/source/services/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routers struct {
	db *gorm.DB
}

func NewRouters(db *gorm.DB) *Routers {
	return &Routers{db: db}
}

func (r *Routers) MountRouters(routeGroup *gin.RouterGroup) {
	auth := routeGroup.Group("/auth")
	auth.POST("/register", register.NewHandler(r.db))
	auth.POST("/login", login.NewHandler(r.db))

	authProtected := routeGroup.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware())
	authProtected.GET("/profile", profile.NewHandler(r.db))
}
