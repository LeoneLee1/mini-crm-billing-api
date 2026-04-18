package productlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) List(ctx context.Context, search, category string, isActive *bool, page, limit int) ([]models.ProductModel, int64, error) {
	query := r.db.WithContext(ctx).Model(&models.ProductModel{})

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", like, like)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var products []models.ProductModel
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
