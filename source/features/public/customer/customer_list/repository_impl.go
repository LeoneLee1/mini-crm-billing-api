package customerlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
	customerrepo "mini-crm-billing-api/source/common/repository/customer_repo"
)

func (r *repositoryImpl) List(ctx context.Context, search, status, createdBy string, page, limit int) ([]models.CustomerModel, int64, error) {
	return customerrepo.GetCustomer(ctx, r.db, search, status, createdBy, page, limit)
}
