package transactionupdatestatus

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

func (r *repositoryImpl) updateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.TransactionModel{}).
		Where("id = ?", id).
		Update("status", status).Error
}
