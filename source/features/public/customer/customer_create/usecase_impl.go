package customercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (u *usecaseImpl) Create(ctx context.Context, createdBy uuid.UUID, req CreateRequest) (*models.CustomerModel, error) {
	status := req.Status
	if status == "" {
		status = "active"
	}

	customer := &models.CustomerModel{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Status:    status,
		CreatedBy: createdBy,
	}

	if err := u.repo.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
