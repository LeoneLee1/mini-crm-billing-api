package productupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type UpdateRequest struct {
	Name        *string
	Description *string
	Price       *float64
	Unit        *string
	Category    *string
	IsActive    *bool
}

type Usecase interface {
	Update(ctx context.Context, id string, req UpdateRequest) (*models.ProductModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
