package profile

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (u *usecaseImpl) GetProfile(ctx context.Context, userID string) (*models.UserModel, error) {
	return u.repo.FindByID(ctx, userID)
}
