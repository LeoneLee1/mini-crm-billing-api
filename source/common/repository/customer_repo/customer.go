package customerrepo

import (
	"context"
	"mini-crm-billing-api/source/common/models"

	"gorm.io/gorm"
)

func FindByID(ctx context.Context, db *gorm.DB, customerID string) (*models.CustomerModel, error) {
	var customer models.CustomerModel
	err := db.WithContext(ctx).Where("id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func GetCustomer(ctx context.Context, db *gorm.DB, search, status, createdBy string, page, limit int) ([]models.CustomerModel, int64, error) {
	query := db.WithContext(ctx).Model(&models.CustomerModel{})

	if createdBy != "" {
		query = query.Where("created_by = ?", createdBy)
	}

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("name LIKE ? OR email LIKE ? OR phone LIKE ?", like, like, like)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var customers []models.CustomerModel
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}
