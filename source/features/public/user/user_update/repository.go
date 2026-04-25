package userupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (*models.UserModel, error)
	IsEmailTakenByOther(ctx context.Context, email, excludeID string) (bool, error)
	UpdateUser(ctx context.Context, user *models.UserModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
