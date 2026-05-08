package profileupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/auth/profile_update/body"
)

var ErrUserNotFound = errors.New("user not found")

type Usecase interface {
	Update(ctx context.Context, userID string, req body.UpdateRequest) (*models.UserModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
