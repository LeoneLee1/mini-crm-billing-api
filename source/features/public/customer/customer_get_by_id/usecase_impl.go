package customergetbyid

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func (u *usecaseImpl) GetByID(ctx context.Context, id string) (*models.CustomerModel, error) {
	customer, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		return nil, err
	}
	return customer, nil
}
