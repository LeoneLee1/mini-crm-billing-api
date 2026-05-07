package login

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	DeleteOldRefreshToken(ctx context.Context, userID uuid.UUID) error
	SaveRefreshToken(ctx context.Context, token *models.RefreshTokenModel) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
