package invoicegetbyid

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (r *repositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error) {
	var invoice models.InvoiceModel
	err := r.db.WithContext(ctx).
		Preload("Transaction").
		Preload("Transaction.Customer").
		Preload("Transaction.Items").
		Preload("Transaction.Items.Product").
		Preload("Payments").
		First(&invoice, "invoices.id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}
