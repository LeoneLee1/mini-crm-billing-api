package userlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (r *repositoryImpl) List(ctx context.Context, search, role string, page, limit int) ([]models.UserModel, int64, error) {
	query := r.db.WithContext(ctx).Model(&models.UserModel{})

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("name LIKE ? OR email LIKE ?", like, like)
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var users []models.UserModel
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
