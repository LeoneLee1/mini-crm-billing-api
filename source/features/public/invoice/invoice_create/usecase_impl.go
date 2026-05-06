package invoicecreate

import (
	"context"
	"errors"
	"time"

	invoicenumberutils "mini-crm-billing-api/source/common/glob_utils/invoice_number_utils"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/invoice/invoice_create/body"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *usecaseImpl) Create(ctx context.Context, req *body.InvoiceCreateRequest) (*models.InvoiceModel, error) {
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		return nil, errors.New("invalid due_date format, use YYYY-MM-DD")
	}

	transactionID, err := uuid.Parse(req.TransactionID)
	if err != nil {
		return nil, errors.New("invalid transaction_id")
	}

	trx, err := u.repo.FindTransactionByID(ctx, transactionID)
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	if trx.Status != string(models.TransactionConfirmed) {
		return nil, errors.New("transaction must be confirmed before creating an invoice")
	}

	existing, err := u.repo.FindByTransactionID(ctx, transactionID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to check existing invoice")
	}
	if existing != nil {
		return nil, errors.New("invoice already exists for this transaction: " + existing.InvoiceNumber)
	}

	invoiceNumber, err := invoicenumberutils.GenerateInvoiceNumber(u.db)
	if err != nil {
		return nil, errors.New("failed to generate invoice number")
	}

	invoice := &models.InvoiceModel{
		TransactionID: transactionID,
		InvoiceNumber: invoiceNumber,
		DueDate:       dueDate,
		Status:        string(models.InvoiceUnPaid),
		Notes:         req.Notes,
	}

	if err := u.repo.Create(ctx, invoice); err != nil {
		return nil, errors.New("failed to save invoice")
	}

	return u.repo.FindByID(ctx, invoice.ID)
}
