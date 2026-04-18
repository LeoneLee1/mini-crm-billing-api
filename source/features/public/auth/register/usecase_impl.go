package register

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/common/models"
	"time"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Register(ctx context.Context, name, email, password string) (*RegisterResponse, error) {
	_, err := u.repo.FindByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := hashingpassword.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.UserModel{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "staff",
	}
	if err := u.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	accessToken, err := jwtutils.CreateAccessToken(user.ID.String(), user.Name, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwtutils.CreateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	tokenModel := &models.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(30 * 24 * time.Hour),
	}
	if err := u.repo.SaveRefreshToken(ctx, tokenModel); err != nil {
		return nil, err
	}

	return &RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}
