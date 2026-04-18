package productdelete

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) FindByID(ctx context.Context, id string) (*models.ProductModel, error) {
	var product models.ProductModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repositoryImpl) DeleteProduct(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.ProductModel{}).Error
}
