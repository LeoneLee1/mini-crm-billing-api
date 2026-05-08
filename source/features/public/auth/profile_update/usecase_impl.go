package profileupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/auth/profile_update/body"
)

// update implements [Usecase].
func (u *usecaseImpl) Update(ctx context.Context, userID string, req body.UpdateRequest) (*models.UserModel, error) {
	existing, err := u.repo.FindByEmail(ctx, req.Email)
	if err == nil && existing.ID.String() != userID {
		return nil, errors.New("email already in use")
	}

	userModel := &models.UserModel{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := u.repo.UpdateProfile(ctx, userID, userModel); err != nil {
		return nil, err
	}

	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
