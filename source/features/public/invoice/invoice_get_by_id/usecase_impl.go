package invoicegetbyid

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *usecaseImpl) getByID(ctx context.Context, id uuid.UUID) (*models.InvoiceModel, error) {
	invoice, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invoice not found")
		}
		return nil, err
	}
	return invoice, nil
}
