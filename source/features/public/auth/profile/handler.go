package profile

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) gin.HandlerFunc {

	handler := Handler{
		db: db,
	}

	return handler.Impl
}
