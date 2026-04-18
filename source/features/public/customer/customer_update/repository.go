package customerupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (*models.CustomerModel, error)
	UpdateCustomer(ctx context.Context, customer *models.CustomerModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
