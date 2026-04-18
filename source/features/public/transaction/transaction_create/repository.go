package transactioncreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	create(ctx context.Context, transaction *models.TransactionModel) error
	customerByID(ctx context.Context, customerID uuid.UUID) error
	productByID(ctx context.Context, productID uuid.UUID) (*models.ProductModel, error)
	transactionByID(ctx context.Context, transactionID uuid.UUID) (*models.TransactionModel, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
