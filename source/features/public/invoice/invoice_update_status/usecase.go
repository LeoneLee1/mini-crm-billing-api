package invoiceupdatestatus

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

type Usecase interface {
	updateStatus(ctx context.Context, id uuid.UUID, status string) (*models.InvoiceModel, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
