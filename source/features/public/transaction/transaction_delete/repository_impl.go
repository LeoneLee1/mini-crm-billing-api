package transactiondelete

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (r *repositoryImpl) transactionByID(ctx context.Context, id uuid.UUID) (*models.TransactionModel, error) {
	var transaction models.TransactionModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *repositoryImpl) deleteTransaction(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.TransactionModel{}).Error
}

func (r *repositoryImpl) deleteItemsByTransactionID(ctx context.Context, transactionID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("transaction_id = ?", transactionID).
		Delete(&models.TransactionItemModel{}).Error
}
