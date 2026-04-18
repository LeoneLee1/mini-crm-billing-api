package login

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/common/models"
	"time"
)

func (u *usecaseImpl) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !hashingpassword.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
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

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}
