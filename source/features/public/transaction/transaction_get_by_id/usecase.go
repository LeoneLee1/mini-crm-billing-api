package transactiongetbyid

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type Usecase interface {
	transactionGetByID(ctx context.Context, transactionID string) (*models.TransactionModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
