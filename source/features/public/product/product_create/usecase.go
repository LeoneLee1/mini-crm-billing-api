package productcreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Name        string
	Description string
	Price       float64
	Unit        string
	Category    string
	IsActive    *bool
}

type Usecase interface {
	Create(ctx context.Context, createdBy uuid.UUID, req CreateRequest) (*models.ProductModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
