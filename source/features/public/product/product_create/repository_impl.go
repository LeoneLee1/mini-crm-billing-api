package productcreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) CreateProduct(ctx context.Context, product *models.ProductModel) error {
	return r.db.WithContext(ctx).Create(product).Error
}
