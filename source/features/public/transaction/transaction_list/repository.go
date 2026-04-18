package transactionlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	list(ctx context.Context, customerID, status string, page, limit int) ([]models.TransactionModel, int64, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
