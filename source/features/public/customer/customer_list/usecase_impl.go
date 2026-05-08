package customerlist

import (
	"context"
	"math"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/customer/customer_list/body"
)

func (u *usecaseImpl) List(ctx context.Context, filter body.ListFilter) (*body.ListResponse, error) {
	var createdBy string
	if filter.RequesterRole == string(models.UserRoleStaff) {
		createdBy = filter.RequesterID
	}

	customers, total, err := u.repo.List(ctx, filter.Search, filter.Status, createdBy, filter.Page, filter.Limit)
	if err != nil {
		return nil, err
	}

	totalPage := int(math.Ceil(float64(total) / float64(filter.Limit)))

	return &body.ListResponse{
		Customer:  customers,
		Total:     total,
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: totalPage,
	}, nil
}
