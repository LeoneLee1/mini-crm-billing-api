package login

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/common/models"
	userrepo "mini-crm-billing-api/source/common/repository/user_repo"
	"mini-crm-billing-api/source/features/public/auth/login/body"
	"mini-crm-billing-api/source/pkg/logger"
	"time"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Login(ctx context.Context, db *gorm.DB, email, password string) (*body.LoginResponse, error) {
	user, err := userrepo.FindByEmail(ctx, db, email)
	if err != nil {
		logger.Error().Err(err).Str("email", email).Msg("FindByEmail failed")
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

	// delete old refresh token
	if err := u.repo.DeleteOldRefreshToken(ctx, user.ID); err != nil {
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

	return &body.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}
