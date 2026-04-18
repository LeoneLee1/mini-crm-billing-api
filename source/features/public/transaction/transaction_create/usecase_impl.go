package transactioncreate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/transaction/transaction_create/body"

	"github.com/google/uuid"
)

// create implements [Usecase].
func (u *usecaseImpl) create(ctx context.Context, req *body.TransactionRequest, createdBy uuid.UUID) (*models.TransactionModel, error) {
	// validate customer
	customerID, err := uuid.Parse(req.CustomerID)
	if err != nil {
		return nil, errors.New("invalid customer id")
	}

	err = u.repo.customerByID(ctx, customerID)
	if err != nil {
		return nil, errors.New("customer not found")
	}

	// build items and calculate subtotal
	var items []models.TransactionItemModel
	var subtotal float64

	for _, itemReq := range req.Items {
		productID, err := uuid.Parse(itemReq.ProductID)
		if err != nil {
			return nil, errors.New("invalid product id: " + itemReq.ProductID)
		}

		product, err := u.repo.productByID(ctx, productID)
		if err != nil {
			return nil, errors.New("product not found: " + itemReq.ProductID)
		}

		if !product.IsActive {
			return nil, errors.New("product is not active: " + product.Name)
		}

		itemSubtotal := product.Price * itemReq.Quantity
		subtotal += itemSubtotal

		items = append(items, models.TransactionItemModel{
			ProductID: productID,
			Quantity:  itemReq.Quantity,
			UnitPrice: product.Price,
			SubTotal:  itemSubtotal,
		})
	}

	// calculate total
	discount := req.Discount
	tax := req.Tax
	total := subtotal - discount + tax

	transaction := &models.TransactionModel{
		CustomerID: customerID,
		CreatedBy:  createdBy,
		Status:     "draft",
		SubTotal:   subtotal,
		Discount:   discount,
		Tax:        tax,
		Total:      total,
		Notes:      req.Notes,
		Items:      items,
	}

	if err := u.repo.create(ctx, transaction); err != nil {
		return nil, err
	}

	return u.repo.transactionByID(ctx, transaction.ID)
}
