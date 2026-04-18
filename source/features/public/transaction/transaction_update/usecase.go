package transactionupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/transaction/transaction_update/body"

	"github.com/google/uuid"
)

type Usecase interface {
	update(ctx context.Context, id uuid.UUID, req *body.UpdateTransactionRequest) (*models.TransactionModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
