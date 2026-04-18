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
	return u.repo.DeleteCustomer(ctx, id)
}
