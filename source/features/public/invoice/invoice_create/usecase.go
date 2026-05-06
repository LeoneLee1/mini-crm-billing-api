package invoicecreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/invoice/invoice_create/body"

	"gorm.io/gorm"
)

type Usecase interface {
	Create(ctx context.Context, req *body.InvoiceCreateRequest) (*models.InvoiceModel, error)
}

type usecaseImpl struct {
	repo Repository
	db   *gorm.DB
}

func injectUsecase(repo Repository, db *gorm.DB) Usecase {
	return &usecaseImpl{repo: repo, db: db}
}
