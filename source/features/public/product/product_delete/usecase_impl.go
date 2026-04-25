package productdelete

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Delete(ctx context.Context, id string) error {
	_, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	used, err := u.repo.IsUsedInTransactions(ctx, id)
	if err != nil {
		return err
	}
	if used {
		return errors.New("product cannot be deleted because it has been used in transactions; consider deactivating it instead")
	}

	return u.repo.DeleteProduct(ctx, id)
}
