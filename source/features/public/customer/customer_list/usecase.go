package customerlist

import (
	"context"
	"mini-crm-billing-api/source/features/public/customer/customer_list/body"
)

type Usecase interface {
	List(ctx context.Context, filter body.ListFilter) (*body.ListResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
