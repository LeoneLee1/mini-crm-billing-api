package profileupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, userID string) (*models.UserModel, error)
	FindByEmail(ctx context.Context, email string) (*models.UserModel, error)
	UpdateProfile(ctx context.Context, userID string, user *models.UserModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
