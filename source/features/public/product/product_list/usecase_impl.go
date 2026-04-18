package productlist

import (
	"context"
	"math"
)

func (u *usecaseImpl) List(ctx context.Context, filter ListFilter) (*ListResponse, error) {
	products, total, err := u.repo.List(ctx, filter.Search, filter.Category, filter.IsActive, filter.Page, filter.Limit)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(filter.Limit)))

	return &ListResponse{
		Product:   products,
		Total:     total,
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: totalPage,
	}, nil
}
