package customerupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type UpdateRequest struct {
	Name    *string
	Email   *string
	Phone   *string
	Address *string
	Status  *string
}

type Usecase interface {
	Update(ctx context.Context, id string, req UpdateRequest) (*models.CustomerModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
