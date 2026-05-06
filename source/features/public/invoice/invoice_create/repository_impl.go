package invoicecreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (r *repositoryImpl) FindTransactionByID(ctx context.Context, id uuid.UUID) (*models.TransactionModel, error) {
	var trx models.TransactionModel
	if err := r.db.WithContext(ctx).First(&trx, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *repositoryImpl) FindByTransactionID(ctx context.Context, transactionID uuid.UUID) (*models.InvoiceModel, error) {
	var invoice models.InvoiceModel
	if err := r.db.WithContext(ctx).First(&invoice, "transaction_id = ?", transactionID).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *repositoryImpl) Create(ctx context.Context, invoice *models.InvoiceModel) error {
	return r.db.WithContext(ctx).Create(invoice).Error
}

func (r *repositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error) {
	var invoice models.InvoiceModel
	if err := r.db.WithContext(ctx).
		Preload("Transaction").
		First(&invoice, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}
