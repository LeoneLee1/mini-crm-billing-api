package body

type UpdateRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Email   string `json:"email"  binding:"omitempty,email"`
	Phone   string `json:"phone"  binding:"omitempty,min=5,max=20"`
	Address string `json:"address"`
	Status  string `json:"status" binding:"omitempty,oneof=active inactive"`
}
