package transactionlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

// list implements [Repository].
func (r *repositoryImpl) list(ctx context.Context, customerID string, status string, page int, limit int) ([]models.TransactionModel, int64, error) {
	query := r.db.WithContext(ctx).Model(&models.TransactionModel{})

	if customerID != "" {
		query = query.Where("customer_id = ?", customerID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var transaction []models.TransactionModel
	offset := (page - 1) * limit
	err := query.
		Preload("Customer").
		Preload("Creator").
		Preload("Items").
		Preload("Items.Product").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&transaction).Error
	if err != nil {
		return nil, 0, err
	}

	return transaction, total, nil
}
