package body

type TransactionItemRequest struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required,gt=0"`
}

type UpdateTransactionRequest struct {
	Discount *float64                 `json:"discount"`
	Tax      *float64                 `json:"tax"`
	Notes    *string                  `json:"notes"`
	Items    []TransactionItemRequest `json:"items"`
}
