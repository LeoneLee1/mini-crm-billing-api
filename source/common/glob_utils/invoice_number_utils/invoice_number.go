package invoicenumberutils

import (
	"fmt"
	"mini-crm-billing-api/source/common/models"
	"time"

	"gorm.io/gorm"
)

// GenerateInvoiceNumber queries the DB to count today's invoices and returns the next number
// in the format INV-YYYYMMDD-XXXX.
func GenerateInvoiceNumber(db *gorm.DB) (string, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var count int64
	if err := db.Model(&models.InvoiceModel{}).
		Where("created_at >= ? AND created_at < ?", startOfDay, endOfDay).
		Count(&count).Error; err != nil {
		return "", err
	}

	return fmt.Sprintf("INV-%s-%04d", now.Format("20060102"), count+1), nil
}
