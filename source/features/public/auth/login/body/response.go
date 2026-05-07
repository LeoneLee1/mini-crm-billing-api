package body

import "mini-crm-billing-api/source/common/models"

type LoginResponse struct {
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	User         *models.UserModel `json:"user"`
}
