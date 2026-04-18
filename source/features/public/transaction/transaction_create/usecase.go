package transactioncreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/transaction/transaction_create/body"

	"github.com/google/uuid"
)

type Usecase interface {
	create(ctx context.Context, req *body.TransactionRequest, createdBy uuid.UUID) (*models.TransactionModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
