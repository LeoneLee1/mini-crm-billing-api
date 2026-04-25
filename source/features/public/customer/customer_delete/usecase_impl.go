package customerdelete

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Delete(ctx context.Context, id string) error {
	_, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("customer not found")
		}
		return err
	}

	hasTransactions, err := u.repo.HasTransactions(ctx, id)
	if err != nil {
		return err
	}
	if hasTransactions {
		return errors.New("customer cannot be deleted because they have existing transactions")
	}

	return u.repo.DeleteCustomer(ctx, id)
}
