package login

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*models.UserModel, error)
	SaveRefreshToken(ctx context.Context, token *models.RefreshTokenModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
