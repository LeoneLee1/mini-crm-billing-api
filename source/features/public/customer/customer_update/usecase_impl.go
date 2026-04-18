package customerupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Update(ctx context.Context, id string, req UpdateRequest) (*models.CustomerModel, error) {
	customer, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}

	if req.Name != nil {
		customer.Name = *req.Name
	}
	if req.Email != nil {
		customer.Email = *req.Email
	}
	if req.Phone != nil {
		customer.Phone = *req.Phone
	}
	if req.Address != nil {
		customer.Address = *req.Address
	}
	if req.Status != nil {
		customer.Status = *req.Status
	}

	if err := u.repo.UpdateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
