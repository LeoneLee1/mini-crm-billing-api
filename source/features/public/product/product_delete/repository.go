package productdelete

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (*models.ProductModel, error)
	IsUsedInTransactions(ctx context.Context, id string) (bool, error)
	DeleteProduct(ctx context.Context, id string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
