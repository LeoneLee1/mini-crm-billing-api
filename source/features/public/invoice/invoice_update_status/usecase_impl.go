package invoiceupdatestatus

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validStatuses = map[string]bool{
	string(models.InvoiceUnPaid):  true,
	string(models.InvoicePaid):    true,
	string(models.InvoiceOverdue): true,
}

func (u *usecaseImpl) updateStatus(ctx context.Context, id uuid.UUID, status string) (*models.InvoiceModel, error) {
	if !validStatuses[status] {
		return nil, errors.New("invalid status, allowed: unpaid, paid, overdue")
	}

	invoice, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invoice not found")
		}
		return nil, err
	}

	if invoice.Status == string(models.InvoicePaid) {
		return nil, errors.New("paid invoice cannot be changed to any other status")
	}

	if err := u.repo.UpdateStatus(ctx, id, status); err != nil {
		return nil, errors.New("failed to update invoice status")
	}

	invoice.Status = status
	return invoice, nil
}
