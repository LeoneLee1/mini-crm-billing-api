package customerdelete

import (
	"context"
	"errors"
)

func (u *usecaseImpl) Delete(ctx context.Context, customerID string) error {
	_, err := u.repo.FindByID(ctx, customerID)
	if err != nil {
		return err
	}

	hasTransactions, err := u.repo.HasTransactions(ctx, customerID)
	if err != nil {
		return err
	}
	if hasTransactions {
		return errors.New("customer cannot be deleted because they have existing transactions")
	}

	return u.repo.DeleteCustomer(ctx, customerID)
}
