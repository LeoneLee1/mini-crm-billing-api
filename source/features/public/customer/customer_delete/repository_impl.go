package customerdelete

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)


func (r *repositoryImpl) FindByID(ctx context.Context, id string) (*models.CustomerModel, error) {
	var customer models.CustomerModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *repositoryImpl) HasTransactions(ctx context.Context, id string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.TransactionModel{}).Where("customer_id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repositoryImpl) DeleteCustomer(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.CustomerModel{}).Error
}
