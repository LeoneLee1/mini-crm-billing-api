package invoicelist

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	list(ctx context.Context, status, customerID string, createdByFilter *uuid.UUID, page, limit int) (*ListResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
