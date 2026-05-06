package invoicedelete

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *usecaseImpl) delete(ctx context.Context, id uuid.UUID) error {
	invoice, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("invoice not found")
		}
		return err
	}

	if invoice.Status != string(models.InvoiceUnPaid) {
		return errors.New("only unpaid invoices can be deleted")
	}

	return u.repo.Delete(ctx, id)
}
