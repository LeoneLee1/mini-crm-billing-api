package register

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) FindByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	var user models.UserModel
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repositoryImpl) CreateUser(ctx context.Context, user *models.UserModel) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *repositoryImpl) SaveRefreshToken(ctx context.Context, token *models.RefreshTokenModel) error {
	return r.db.WithContext(ctx).Create(token).Error
}
