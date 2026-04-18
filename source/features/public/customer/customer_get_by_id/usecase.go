package customergetbyid

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type Usecase interface {
	GetByID(ctx context.Context, id string) (*models.CustomerModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
