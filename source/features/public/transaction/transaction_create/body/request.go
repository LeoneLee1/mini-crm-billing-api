package body

type TransactionItemRequest struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  float64 `json:"quantity" binding:"required,gt=0"`
}

type TransactionRequest struct {
	CustomerID string                   `json:"customer_id" binding:"required"`
	Items      []TransactionItemRequest `json:"items" binding:"required,min=1"`
	Discount   float64                  `json:"discount"`
	Tax        float64                  `json:"tax"`
	Notes      string                   `json:"notes"`
}
