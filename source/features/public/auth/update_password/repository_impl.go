package updatepassword

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

// findByID implements [Repository].
func (r *repositoryImpl) findByID(ctx context.Context, id string) (*models.UserModel, error) {
	var user models.UserModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// updatePassword implements [Repository].
func (r *repositoryImpl) updatePassword(ctx context.Context, id string, newPassword string) error {
	return r.db.WithContext(ctx).Model(&models.UserModel{}).Where("id = ?", id).Update("password", newPassword).Error
}
