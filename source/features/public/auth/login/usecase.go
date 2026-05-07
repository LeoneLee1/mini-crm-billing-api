package login

import (
	"context"
	"mini-crm-billing-api/source/features/public/auth/login/body"

	"gorm.io/gorm"
)

type Usecase interface {
	Login(ctx context.Context, db *gorm.DB, email, password string) (*body.LoginResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
