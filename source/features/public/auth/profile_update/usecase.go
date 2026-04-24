package profileupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type Usecase interface {
	update(ctx context.Context, id string, req updateRequest) (*models.UserModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}

type updateRequest struct {
	Name     string `json:"name"`
	Email    string `jsom:"email"`
	Password string `json:"password"`
}
