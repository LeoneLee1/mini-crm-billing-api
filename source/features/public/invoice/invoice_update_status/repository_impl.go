package invoiceupdatestatus

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

func (r *repositoryImpl) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.InvoiceModel{}).
		Where("id = ?", id).
		Update("status", status).Error
}
