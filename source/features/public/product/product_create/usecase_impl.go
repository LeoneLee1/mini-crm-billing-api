package productcreate

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"github.com/google/uuid"
)

func (u *usecaseImpl) Create(ctx context.Context, createdBy uuid.UUID, req CreateRequest) (*models.ProductModel, error) {
	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	product := &models.ProductModel{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Unit:        req.Unit,
		Category:    req.Category,
		IsActive:    isActive,
		CreatedBy:   createdBy,
	}

	if err := u.repo.CreateProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
