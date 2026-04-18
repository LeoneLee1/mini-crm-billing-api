package profile

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type Usecase interface {
	GetProfile(ctx context.Context, userID string) (*models.UserModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
