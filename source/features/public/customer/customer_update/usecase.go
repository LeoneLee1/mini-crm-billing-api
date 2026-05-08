package customerupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/customer/customer_update/body"
)

var ErrCustomerNotFound = errors.New("Customer not found")
var ErrForbidden = errors.New("forbidden")

type Usecase interface {
	Update(ctx context.Context, id string, req body.UpdateRequest, requesterID string, requesterRole string) (*models.CustomerModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
