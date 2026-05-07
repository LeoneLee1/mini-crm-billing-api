package login

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

// DeleteOldRefreshToken implements [Repository].
func (r *repositoryImpl) DeleteOldRefreshToken(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.RefreshTokenModel{}).Error
}

func (r *repositoryImpl) SaveRefreshToken(ctx context.Context, token *models.RefreshTokenModel) error {
	return r.db.WithContext(ctx).Create(token).Error
}
