package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentModel struct {
	ID                    uuid.UUID  `gorm:"primaryKey;type:char(36)" json:"id"`
	InvoiceID             uuid.UUID  `gorm:"type:char(36);not null" json:"invoice_id"`
	Amount                float64    `gorm:"type:decimal(15,2);not null" json:"amount"`
	Method                string     `gorm:"varchar(50)" json:"method"`
	MidtransOrderID       string     `gorm:"varchar(100);uniqueIndex" json:"midtrans_order_id"`
	MidtransTransactionID string     `gorm:"varchar(100)" json:"midtrans_transaction_id"`
	MidtransStatus        string     `gorm:"varchar(50)" json:"midtrans_status"`
	PaidAt                *time.Time `json:"paid_at"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Invoice *InvoiceModel `gorm:"foreignKey:InvoiceID" json:"invoice,omitempty"`
}

func (PaymentModel) TableName() string {
	return "payments"
}

func (pm *PaymentModel) BeforeCreate(tx *gorm.DB) (err error) {
	if pm.ID == uuid.Nil {
		pm.ID = uuid.New()
	}

	return
}
