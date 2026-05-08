package customercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/customer/customer_create/body"
)

type Usecase interface {
	Create(ctx context.Context, createdBy string, req body.CreateRequest) (*models.CustomerModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
