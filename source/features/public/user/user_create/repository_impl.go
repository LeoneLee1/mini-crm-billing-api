package usercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.UserModel{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repositoryImpl) CreateUser(ctx context.Context, user *models.UserModel) error {
	return r.db.WithContext(ctx).Create(user).Error
}
