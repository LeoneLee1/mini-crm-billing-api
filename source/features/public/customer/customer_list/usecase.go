package customerlist

import (
	"context"
	"mini-crm-billing-api/source/common/models"
)

type ListFilter struct {
	Search string
	Status string
	Page   int
	Limit  int
}

type ListResponse struct {
	Customer  []models.CustomerModel `json:"customers"`
	Total     int64                  `json:"total"`
	Page      int                    `json:"page"`
	Limit     int                    `json:"limit"`
	TotalPage int                    `json:"total_page"`
}

type Usecase interface {
	List(ctx context.Context, filter ListFilter) (*ListResponse, error)
}

type usecaseImpl struct {
	repo Repository
}

func injectUsecase(repo Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
