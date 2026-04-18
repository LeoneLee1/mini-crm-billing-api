package customercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

type CreateRequest struct {
	Name    string
	Email   string
	Phone   string
	Address string
	Status  string
}

type Usecase interface {
	Create(ctx context.Context, createdBy uuid.UUID, req CreateRequest) (*models.CustomerModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
