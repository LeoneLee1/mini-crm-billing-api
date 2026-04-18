package transactiongetbyid

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

// transactionGetByID implements [Repository].
func (r *repositoryImpl) transactionGetByID(ctx context.Context, transactionID string) (*models.TransactionModel, error) {
	var transaction models.TransactionModel
	err := r.db.
		Preload("Customer").
		Preload("Creator").
		Preload("Items").
		Preload("Items.Product").
		Preload("Invoice").
		Where("id = ?", transactionID).
		First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
