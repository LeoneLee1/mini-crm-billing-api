package customercreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) CreateCustomer(ctx context.Context, customer *models.CustomerModel) error {
	return r.db.WithContext(ctx).Create(customer).Error
}
