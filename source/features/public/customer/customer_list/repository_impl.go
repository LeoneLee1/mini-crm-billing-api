package customerlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) List(ctx context.Context, search, status string, page, limit int) ([]models.CustomerModel, int64, error) {
	query := r.db.WithContext(ctx).Model(&models.CustomerModel{})

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("name LIKE ? OR email LIKE ? OR phone LIKE ?", like, like, like)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var customers []models.CustomerModel
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}
