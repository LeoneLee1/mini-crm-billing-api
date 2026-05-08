package customergetbyid

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(db *gorm.DB) gin.HandlerFunc {
	repo := injectRepository(db)
	usecase := injectUsecase(repo)

	handler := Handler{
		usecase: usecase,
	}

	return handler.Impl
}
