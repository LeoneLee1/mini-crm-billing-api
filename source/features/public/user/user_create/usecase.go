package usercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type CreateRequest struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type Usecase interface {
	Create(ctx context.Context, req CreateRequest) (*models.UserModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
