package userdelete

import "context"

type Usecase interface {
	Delete(ctx context.Context, id string) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
