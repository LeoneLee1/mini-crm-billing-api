package customergetbyid

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	customerrepo "mini-crm-billing-api/source/common/repository/customer_repo"

	"gorm.io/gorm"
)

func (r *repositoryImpl) FindByID(ctx context.Context, customerID string) (*models.CustomerModel, error) {
	customer, err := customerrepo.FindByID(ctx, r.db, customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, err
	}
	return customer, nil
}
