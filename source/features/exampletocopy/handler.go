package exampletocopy

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
	repo Repository
	usecase Usecase
}

func NewHandler(db *gorm.DB) gin.HandlerFunc {
	repo := injectRepository(db)
	usecase := injectUsecase(repo)

	handler := Handler{
		db: db, 
		repo: repo, 
		usecase: usecase,
	}

	return handler.Impl
}