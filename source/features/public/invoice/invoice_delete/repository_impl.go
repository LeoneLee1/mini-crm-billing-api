package invoicedelete

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (r *repositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error) {
	var invoice models.InvoiceModel
	if err := r.db.WithContext(ctx).First(&invoice, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *repositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.InvoiceModel{}, "id = ?", id).Error
}
