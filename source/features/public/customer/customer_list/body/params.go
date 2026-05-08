package body

type ListFilter struct {
	Search        string
	Status        string
	Page          int
	Limit         int
	RequesterID   string
	RequesterRole string
}
