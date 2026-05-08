package customercreate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/customer/customer_create/body"

	"github.com/google/uuid"
)

func (u *usecaseImpl) Create(ctx context.Context, createdBy string, req body.CreateRequest) (*models.CustomerModel, error) {
	createdByUUID, err := uuid.Parse(createdBy)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	customer := &models.CustomerModel{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Status:    string(models.CustomerStatusActive),
		CreatedBy: createdByUUID,
	}

	if err := u.repo.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
