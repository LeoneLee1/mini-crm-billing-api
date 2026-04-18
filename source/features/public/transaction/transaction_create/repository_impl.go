package transactioncreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

// create implements [Repository].
func (r *repositoryImpl) create(ctx context.Context, transaction *models.TransactionModel) error {
	return r.db.WithContext(ctx).Create(&transaction).Error
}

// customerByID implements [Repository].
func (r *repositoryImpl) customerByID(ctx context.Context, customerID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?", customerID).First(&models.CustomerModel{}).Error
}

// productByID implements [Repository].
func (r *repositoryImpl) productByID(ctx context.Context, productID uuid.UUID) (*models.ProductModel, error) {
	var product models.ProductModel
	err := r.db.WithContext(ctx).Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

// transactionByID implements [Repository].
func (r *repositoryImpl) transactionByID(ctx context.Context, transactionID uuid.UUID) (*models.TransactionModel, error) {
	var transaction models.TransactionModel
	err := r.db.WithContext(ctx).Where("id = ?", transactionID).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
