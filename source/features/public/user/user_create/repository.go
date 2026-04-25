package usercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	IsEmailTaken(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.UserModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
