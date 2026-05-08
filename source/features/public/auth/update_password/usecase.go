package updatepassword

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/features/public/auth/update_password/body"
)

var ErrUserNotFound = errors.New("user not found")

type Usecase interface {
	Update(ctx context.Context, userID string, req body.UpdateRequest) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
