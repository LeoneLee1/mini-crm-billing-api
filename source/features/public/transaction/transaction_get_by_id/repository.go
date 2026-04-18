package transactiongetbyid

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	transactionGetByID(ctx context.Context, transactionID string) (*models.TransactionModel, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
