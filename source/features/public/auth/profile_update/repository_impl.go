package profileupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	userrepo "mini-crm-billing-api/source/common/repository/user_repo"

	"gorm.io/gorm"
)

// FindByEmail implements [Repository].
func (r *repositoryImpl) FindByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	return userrepo.FindByEmail(ctx, r.db, email)
}

// FindByID implements [Repository].
func (r *repositoryImpl) FindByID(ctx context.Context, userID string) (*models.UserModel, error) {
	user, err := userrepo.FindByID(ctx, r.db, userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return user, err
}

// UpdateProfile implements [Repository].
func (r *repositoryImpl) UpdateProfile(ctx context.Context, userID string, user *models.UserModel) error {
	return r.db.WithContext(ctx).Where("id = ?", userID).Updates(user).Error
}
