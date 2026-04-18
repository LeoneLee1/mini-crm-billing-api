package transactiongetbyid

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
)

// transactionGetByID implements [Usecase].
func (u *usecaseImpl) transactionGetByID(ctx context.Context, transactionID string) (*models.TransactionModel, error) {
	transaction, err := u.repo.transactionGetByID(ctx, transactionID)
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}
