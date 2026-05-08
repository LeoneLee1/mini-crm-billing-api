package customerdelete

import (
	"context"
	"errors"
)

var ErrCustomerNotFound = errors.New("customer not found")

type Usecase interface {
	Delete(ctx context.Context, customerID string) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
