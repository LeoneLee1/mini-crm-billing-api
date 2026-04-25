package userupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Update(ctx context.Context, id string, req UpdateRequest) (*models.UserModel, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		taken, err := u.repo.IsEmailTakenByOther(ctx, *req.Email, id)
		if err != nil {
			return nil, err
		}
		if taken {
			return nil, errors.New("email already in use")
		}
		user.Email = *req.Email
	}
	if req.Role != nil {
		user.Role = *req.Role
	}

	if err := u.repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
