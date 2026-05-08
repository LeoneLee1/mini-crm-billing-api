package customergetbyid

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
)

var ErrCustomerNotFound = errors.New("Customer not found")

type Usecase interface {
	GetByID(ctx context.Context, customerID string) (*models.CustomerModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
