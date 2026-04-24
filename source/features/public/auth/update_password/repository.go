package updatepassword

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	findByID(ctx context.Context, id string) (*models.UserModel, error)
	updatePassword(ctx context.Context, id string, newPassword string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
