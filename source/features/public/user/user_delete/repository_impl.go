package userdelete

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

func (r *repositoryImpl) DeleteUser(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.UserModel{}).Error
}
