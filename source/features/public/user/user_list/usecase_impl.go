package userlist

import (
	"context"
	"math"
)

func (u *usecaseImpl) List(ctx context.Context, filter ListFilter) (*ListResponse, error) {
	users, total, err := u.repo.List(ctx, filter.Search, filter.Role, filter.Page, filter.Limit)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(filter.Limit)))

	return &ListResponse{
		Users:     users,
		Total:     total,
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: totalPage,
	}, nil
}
