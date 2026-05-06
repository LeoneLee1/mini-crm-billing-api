package invoicedelete

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
