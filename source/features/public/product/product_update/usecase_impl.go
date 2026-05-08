package productupdate

import (
	"context"
	"errors"
	productcategoryutils "mini-crm-billing-api/source/common/glob_utils/product_category_utils"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func (u *usecaseImpl) Update(ctx context.Context, id string, req UpdateRequest) (*models.ProductModel, error) {
	product, err := u.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if req.Category != nil && !productcategoryutils.IsValidCategory(*req.Category) {
		return nil, errors.New("Invalid category")
	}

	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Unit != nil {
		product.Unit = *req.Unit
	}
	if req.Category != nil {
		product.Category = productcategoryutils.NormalizeCategory(*req.Category)
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}

	if err := u.repo.UpdateProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
