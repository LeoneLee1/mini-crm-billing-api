package services

import (
	"mini-crm-billing-api/source/features/public/auth/login"
	"mini-crm-billing-api/source/features/public/auth/profile"
	"mini-crm-billing-api/source/features/public/auth/register"
	customercreate "mini-crm-billing-api/source/features/public/customer/customer_create"
	customerdelete "mini-crm-billing-api/source/features/public/customer/customer_delete"
	customergetbyid "mini-crm-billing-api/source/features/public/customer/customer_get_by_id"
	customerlist "mini-crm-billing-api/source/features/public/customer/customer_list"
	customerupdate "mini-crm-billing-api/source/features/public/customer/customer_update"
	productcreate "mini-crm-billing-api/source/features/public/product/product_create"
	productdelete "mini-crm-billing-api/source/features/public/product/product_delete"
	productgetbyid "mini-crm-billing-api/source/features/public/product/product_get_by_id"
	productlist "mini-crm-billing-api/source/features/public/product/product_list"
	productupdate "mini-crm-billing-api/source/features/public/product/product_update"
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
	// Auth public
	auth := routeGroup.Group("/auth")
	auth.POST("/register", register.NewHandler(r.db))
	auth.POST("/login", login.NewHandler(r.db))

	// Auth protected
	authProtected := routeGroup.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware())
	authProtected.GET("/profile", profile.NewHandler(r.db))

	// Customer (protected)
	customer := routeGroup.Group("/customers")
	customer.Use(middleware.AuthMiddleware())
	customer.GET("", customerlist.NewHandler(r.db))
	customer.POST("", customercreate.NewHandler(r.db))
	customer.GET("/:id", customergetbyid.NewHandler(r.db))
	customer.PUT("/:id", customerupdate.NewHandler(r.db))
	customer.DELETE("/:id", customerdelete.NewHandler(r.db))

	// Product (protected)
	product := routeGroup.Group("/products")
	product.Use(middleware.AuthMiddleware())
	product.GET("", productlist.NewHandler(r.db))
	product.POST("", productcreate.NewHandler(r.db))
	product.GET("/:id", productgetbyid.NewHandler(r.db))
	product.PUT("/:id", productupdate.NewHandler(r.db))
	product.DELETE("/:id", productdelete.NewHandler(r.db))
}
