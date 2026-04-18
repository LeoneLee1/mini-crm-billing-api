package transactionupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	transactionByID(ctx context.Context, id uuid.UUID) (*models.TransactionModel, error)
	productByID(ctx context.Context, id uuid.UUID) (*models.ProductModel, error)
	updateTransaction(ctx context.Context, transaction *models.TransactionModel) error
	deleteItemsByTransactionID(ctx context.Context, transactionID uuid.UUID) error
	createItems(ctx context.Context, items []models.TransactionItemModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
