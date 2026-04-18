package transactionupdatestatus

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	updateStatus(ctx context.Context, id uuid.UUID, status string) error
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
