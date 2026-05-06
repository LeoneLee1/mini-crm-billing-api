package body

type InvoiceCreateRequest struct {
	TransactionID string `json:"transaction_id" binding:"required"`
	DueDate       string `json:"due_date" binding:"required"`
	Notes         string `json:"notes"`
}
