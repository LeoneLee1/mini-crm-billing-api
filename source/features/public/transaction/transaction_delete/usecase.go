package transactiondelete

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	delete(ctx context.Context, id uuid.UUID) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
