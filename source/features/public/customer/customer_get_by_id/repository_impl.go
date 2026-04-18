package customergetbyid

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
