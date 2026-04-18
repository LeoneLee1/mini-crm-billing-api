package login

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type LoginResponse struct {
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	User         *models.UserModel `json:"user"`
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
