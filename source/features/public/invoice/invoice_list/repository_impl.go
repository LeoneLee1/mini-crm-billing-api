package invoicelist

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (r *repositoryImpl) list(ctx context.Context, status, customerID string, createdByFilter *uuid.UUID, page, limit int) ([]models.InvoiceModel, int64, error) {
	query := r.db.WithContext(ctx).Model(&models.InvoiceModel{}).
		Joins("JOIN transactions ON transactions.id = invoices.transaction_id AND transactions.deleted_at IS NULL")

	if status != "" {
		query = query.Where("invoices.status = ?", status)
	}

	if customerID != "" {
		query = query.Where("transactions.customer_id = ?", customerID)
	}

	if createdByFilter != nil {
		query = query.Where("transactions.created_by = ?", *createdByFilter)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var invoices []models.InvoiceModel
	offset := (page - 1) * limit
	err := query.
		Preload("Transaction").
		Preload("Transaction.Customer").
		Preload("Payments").
		Order("invoices.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&invoices).Error
	if err != nil {
		return nil, 0, err
	}

	return invoices, total, nil
}
