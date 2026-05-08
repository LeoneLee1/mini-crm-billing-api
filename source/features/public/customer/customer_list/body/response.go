package body

import "mini-crm-billing-api/source/common/models"

type ListResponse struct {
	Customer  []models.CustomerModel `json:"customers"`
	Total     int64                  `json:"total"`
	Page      int                    `json:"page"`
	Limit     int                    `json:"limit"`
	TotalPage int                    `json:"total_page"`
}
