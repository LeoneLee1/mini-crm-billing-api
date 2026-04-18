package transactionupdate

import (
	"context"
	"errors"
	"mini-crm-billing-api/source/common/models"
	"mini-crm-billing-api/source/features/public/transaction/transaction_update/body"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (u *usecaseImpl) update(ctx context.Context, id uuid.UUID, req *body.UpdateTransactionRequest) (*models.TransactionModel, error) {
	transaction, err := u.repo.transactionByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}

	if transaction.Status != "draft" {
		return nil, errors.New("only draft transactions can be edited")
	}

	if req.Notes != nil {
		transaction.Notes = *req.Notes
	}

	if len(req.Items) > 0 {
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
				TransactionID: transaction.ID,
				ProductID:     productID,
				Quantity:      itemReq.Quantity,
				UnitPrice:     product.Price,
				SubTotal:      itemSubtotal,
			})
		}

		if err := u.repo.deleteItemsByTransactionID(ctx, transaction.ID); err != nil {
			return nil, err
		}
		if err := u.repo.createItems(ctx, items); err != nil {
			return nil, err
		}

		transaction.SubTotal = subtotal
	}

	if req.Discount != nil {
		transaction.Discount = *req.Discount
	}
	if req.Tax != nil {
		transaction.Tax = *req.Tax
	}

	if len(req.Items) > 0 || req.Discount != nil || req.Tax != nil {
		transaction.Total = transaction.SubTotal - transaction.Discount + transaction.Tax
	}

	if err := u.repo.updateTransaction(ctx, transaction); err != nil {
		return nil, err
	}

	return u.repo.transactionByID(ctx, transaction.ID)
}
