package userrepo

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func FindByEmail(ctx context.Context, db *gorm.DB, email string) (*models.UserModel, error) {
	var user models.UserModel
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
