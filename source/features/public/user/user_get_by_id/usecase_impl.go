package usergetbyid

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func (u *usecaseImpl) GetByID(ctx context.Context, id string) (*models.UserModel, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
