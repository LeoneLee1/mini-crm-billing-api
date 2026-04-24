package profileupdate

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

// update implements [Usecase].
func (u *usecaseImpl) update(ctx context.Context, id string, req updateRequest) (*models.UserModel, error) {
	user, err := u.repo.findByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	hashedPassword, err := hashingpassword.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	userModel := &models.UserModel{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := u.repo.updateProfile(ctx, id, userModel); err != nil {
		return nil, err
	}

	return user, nil
}
