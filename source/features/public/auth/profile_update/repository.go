package profileupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	findByID(ctx context.Context, id string) (*models.UserModel, error)
	updateProfile(ctx context.Context, id string, profile *models.UserModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
