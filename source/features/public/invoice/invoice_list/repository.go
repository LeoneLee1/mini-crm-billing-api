package invoicelist

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ListResponse struct {
	Invoices  []models.InvoiceModel `json:"invoices"`
	Total     int64                 `json:"total"`
	Page      int                   `json:"page"`
	Limit     int                   `json:"limit"`
	TotalPage int                   `json:"total_page"`
}

type Repository interface {
	list(ctx context.Context, status, customerID string, createdByFilter *uuid.UUID, page, limit int) ([]models.InvoiceModel, int64, error)
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
