package register

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type RegisterResponse struct {
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	User         *models.UserModel `json:"user"`
}

type Usecase interface {
	Register(ctx context.Context, name, email, password string) (*RegisterResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
