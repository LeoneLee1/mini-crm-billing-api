package invoicecreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindTransactionByID(ctx context.Context, id uuid.UUID) (*models.TransactionModel, error)
	FindByTransactionID(ctx context.Context, transactionID uuid.UUID) (*models.InvoiceModel, error)
	Create(ctx context.Context, invoice *models.InvoiceModel) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
