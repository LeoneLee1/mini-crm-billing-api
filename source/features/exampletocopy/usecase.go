package exampletocopy

import "context"

type Usecase interface {
	LogUserAccess(ctx context.Context, userID uint, requesterIP string) error
}

type usecaseImpl struct{
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}