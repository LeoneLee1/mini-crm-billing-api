package userupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) FindByID(ctx context.Context, id string) (*models.UserModel, error) {
	var user models.UserModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) IsEmailTakenByOther(ctx context.Context, email, excludeID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.UserModel{}).
		Where("email = ? AND id != ?", email, excludeID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repositoryImpl) UpdateUser(ctx context.Context, user *models.UserModel) error {
	return r.db.WithContext(ctx).Save(user).Error
}
