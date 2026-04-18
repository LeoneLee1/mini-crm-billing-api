package transactionupdatestatus

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var allowedTransitions = map[string]map[string]bool{
	"draft":     {"confirmed": true, "cancelled": true},
	"confirmed": {"paid": true, "cancelled": true},
}

func (u *usecaseImpl) updateStatus(ctx context.Context, id uuid.UUID, status string) error {
	transaction, err := u.repo.transactionByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("transaction not found")
		}
		return err
	}

	transitions, ok := allowedTransitions[transaction.Status]
	if !ok {
		return errors.New("transaction status '" + transaction.Status + "' cannot be changed")
	}

	if !transitions[status] {
		return errors.New("cannot transition from '" + transaction.Status + "' to '" + status + "'")
	}

	return u.repo.updateStatus(ctx, id, status)
}
