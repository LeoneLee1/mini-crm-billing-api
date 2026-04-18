package transactionlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type ListResponse struct {
	Transaction []models.TransactionModel `json:"transaction"`
	Total       int64                     `json:"total"`
	Page        int                       `json:"page"`
	Limit       int                       `json:"limit"`
	TotalPage   int                       `json:"total_page"`
}

type Usecase interface {
	list(ctx context.Context, customerID, status string, page, limit int) (*ListResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
