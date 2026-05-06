package services

import (
	"mini-crm-billing-api/source/features/public/auth/login"
	"mini-crm-billing-api/source/features/public/auth/profile"
	profileupdate "mini-crm-billing-api/source/features/public/auth/profile_update"
	updatepassword "mini-crm-billing-api/source/features/public/auth/update_password"
	customercreate "mini-crm-billing-api/source/features/public/customer/customer_create"
	customerdelete "mini-crm-billing-api/source/features/public/customer/customer_delete"
	customergetbyid "mini-crm-billing-api/source/features/public/customer/customer_get_by_id"
	customerlist "mini-crm-billing-api/source/features/public/customer/customer_list"
	customerupdate "mini-crm-billing-api/source/features/public/customer/customer_update"
	invoicecreate "mini-crm-billing-api/source/features/public/invoice/invoice_create"
	invoicedelete "mini-crm-billing-api/source/features/public/invoice/invoice_delete"
	invoicegetbyid "mini-crm-billing-api/source/features/public/invoice/invoice_get_by_id"
	invoicelist "mini-crm-billing-api/source/features/public/invoice/invoice_list"
	invoiceupdatestatus "mini-crm-billing-api/source/features/public/invoice/invoice_update_status"
	productcreate "mini-crm-billing-api/source/features/public/product/product_create"
	productdelete "mini-crm-billing-api/source/features/public/product/product_delete"
	productgetbyid "mini-crm-billing-api/source/features/public/product/product_get_by_id"
	productlist "mini-crm-billing-api/source/features/public/product/product_list"
	productupdate "mini-crm-billing-api/source/features/public/product/product_update"
	transactioncreate "mini-crm-billing-api/source/features/public/transaction/transaction_create"
	transactiondelete "mini-crm-billing-api/source/features/public/transaction/transaction_delete"
	transactionlist "mini-crm-billing-api/source/features/public/transaction/transaction_list"
	transactionupdate "mini-crm-billing-api/source/features/public/transaction/transaction_update"
	transactionupdatestatus "mini-crm-billing-api/source/features/public/transaction/transaction_update_status"
	usercreate "mini-crm-billing-api/source/features/public/user/user_create"
	userdelete "mini-crm-billing-api/source/features/public/user/user_delete"
	usergetbyid "mini-crm-billing-api/source/features/public/user/user_get_by_id"
	userlist "mini-crm-billing-api/source/features/public/user/user_list"
	userupdate "mini-crm-billing-api/source/features/public/user/user_update"
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
	auth.POST("/login", login.NewHandler(r.db))

	// Auth protected
	authProtected := routeGroup.Group("/auth")
	authProtected.Use(middleware.AuthMiddleware())
	authProtected.GET("/profile", profile.NewHandler(r.db))
	authProtected.PUT("/profile/update", profileupdate.NewHandler(r.db))
	authProtected.PUT("/profile/update-password", updatepassword.NewHandler(r.db))

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

	// Transaction (protected)
	transaction := routeGroup.Group("/transactions")
	transaction.Use(middleware.AuthMiddleware())
	transaction.GET("", transactionlist.NewHandler(r.db))
	transaction.POST("", transactioncreate.NewHandler(r.db))
	transaction.PUT("/:id", transactionupdate.NewHandler(r.db))
	transaction.PATCH("/:id/status", transactionupdatestatus.NewHandler(r.db))
	transaction.DELETE("/:id", transactiondelete.NewHandler(r.db))

	// Invoice (protected)
	invoice := routeGroup.Group("/invoice")
	invoice.Use(middleware.AuthMiddleware())
	invoice.GET("", invoicelist.NewHandler(r.db))
	invoice.POST("", invoicecreate.NewHandler(r.db))
	invoice.GET("/:id", invoicegetbyid.NewHandler(r.db))
	invoice.PATCH("/:id/status", invoiceupdatestatus.NewHandler(r.db))
	invoice.DELETE("/:id", invoicedelete.NewHandler(r.db))

	// User management (admin only)
	user := routeGroup.Group("/users")
	user.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	user.GET("", userlist.NewHandler(r.db))
	user.POST("", usercreate.NewHandler(r.db))
	user.GET("/:id", usergetbyid.NewHandler(r.db))
	user.PUT("/:id", userupdate.NewHandler(r.db))
	user.DELETE("/:id", userdelete.NewHandler(r.db))
}
