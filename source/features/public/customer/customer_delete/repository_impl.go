package customerdelete

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

func (r *repositoryImpl) HasTransactions(ctx context.Context, customerID string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.TransactionModel{}).Where("customer_id = ?", customerID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repositoryImpl) DeleteCustomer(ctx context.Context, customerID string) error {
	return r.db.WithContext(ctx).Where("id = ?", customerID).Delete(&models.CustomerModel{}).Error
}
