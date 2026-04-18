package transactionupdatestatus

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	transactionByID(ctx context.Context, id uuid.UUID) (*models.TransactionModel, error)
	updateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
