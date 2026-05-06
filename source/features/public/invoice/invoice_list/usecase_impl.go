package invoicelist

import (
	"context"
	"math"

	"github.com/google/uuid"
)

func (u *usecaseImpl) list(ctx context.Context, status, customerID string, createdByFilter *uuid.UUID, page, limit int) (*ListResponse, error) {
	invoices, total, err := u.repo.list(ctx, status, customerID, createdByFilter, page, limit)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	return &ListResponse{
		Invoices:  invoices,
		Total:     total,
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
	}, nil
}
