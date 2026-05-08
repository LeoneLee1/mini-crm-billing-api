package customerupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/customer/customer_update/body"
)

func (u *usecaseImpl) Update(ctx context.Context, id string, req body.UpdateRequest, requesterID string, requesterRole string) (*models.CustomerModel, error) {
	customer, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if requesterRole == string(models.UserRoleStaff) && customer.CreatedBy.String() != requesterID {
		return nil, ErrForbidden
	}

	customer.Name    = req.Name
	customer.Email   = req.Email
	customer.Phone   = req.Phone
	customer.Address = req.Address
	if req.Status != "" {
		customer.Status = req.Status
	}

	if err := u.repo.UpdateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
