package transactionlist

import (
	"context"
	"math"
)

// list implements [Usecase].
func (u *usecaseImpl) list(ctx context.Context, customerID string, status string, page int, limit int) (*ListResponse, error) {
	transaction, total, err := u.repo.list(ctx, customerID, status, page, limit)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	return &ListResponse{
		Transaction: transaction,
		Total:       total,
		Page:        page,
		Limit:       limit,
		TotalPage:   totalPage,
	}, nil
}
