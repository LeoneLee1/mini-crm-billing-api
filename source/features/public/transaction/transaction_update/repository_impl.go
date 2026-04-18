package transactionupdate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (r *repositoryImpl) transactionByID(ctx context.Context, id uuid.UUID) (*models.TransactionModel, error) {
	var transaction models.TransactionModel
	err := r.db.WithContext(ctx).
		Preload("Items.Product").
		Preload("Customer").
		Where("id = ?", id).
		First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *repositoryImpl) productByID(ctx context.Context, id uuid.UUID) (*models.ProductModel, error) {
	var product models.ProductModel
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repositoryImpl) updateTransaction(ctx context.Context, transaction *models.TransactionModel) error {
	return r.db.WithContext(ctx).Save(transaction).Error
}

func (r *repositoryImpl) deleteItemsByTransactionID(ctx context.Context, transactionID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("transaction_id = ?", transactionID).
		Delete(&models.TransactionItemModel{}).Error
}

func (r *repositoryImpl) createItems(ctx context.Context, items []models.TransactionItemModel) error {
	return r.db.WithContext(ctx).Create(&items).Error
}
