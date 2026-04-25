package usercreate

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"
	"mini-crm-billing-api/source/common/models"
)

func (u *usecaseImpl) Create(ctx context.Context, req CreateRequest) (*models.UserModel, error) {
	taken, err := u.repo.IsEmailTaken(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if taken {
		return nil, errors.New("email already in use")
	}

	hashed, err := hashingpassword.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.UserModel{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		Role:     req.Role,
	}

	if err := u.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
