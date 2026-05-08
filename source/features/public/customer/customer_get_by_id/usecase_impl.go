package customergetbyid

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

func (u *usecaseImpl) GetByID(ctx context.Context, customerID string) (*models.CustomerModel, error) {
	return u.repo.FindByID(ctx, customerID)
}
