package invoiceupdatestatus

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
