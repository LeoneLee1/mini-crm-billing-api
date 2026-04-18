package transactiondelete

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *usecaseImpl) delete(ctx context.Context, id uuid.UUID) error {
	transaction, err := u.repo.transactionByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("transaction not found")
		}
		return err
	}

	if transaction.Status != "draft" {
		return errors.New("only draft transactions can be deleted")
	}

	if err := u.repo.deleteItemsByTransactionID(ctx, id); err != nil {
		return err
	}

	return u.repo.deleteTransaction(ctx, id)
}
