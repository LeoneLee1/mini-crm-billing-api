package updatepassword

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	userrepo "mini-crm-billing-api/source/common/repository/user_repo"

	"gorm.io/gorm"
)

// findByID implements [Repository].
func (r *repositoryImpl) FindByID(ctx context.Context, userID string) (*models.UserModel, error) {
	user, err := userrepo.FindByID(ctx, r.db, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return user, err
}

// updatePassword implements [Repository].
func (r *repositoryImpl) UpdatePassword(ctx context.Context, userID string, newPassword string) error {
	return r.db.WithContext(ctx).Model(&models.UserModel{}).Where("id = ?", userID).Update("password", newPassword).Error
}
