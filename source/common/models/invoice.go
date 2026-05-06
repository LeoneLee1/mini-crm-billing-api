package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceStatus string

const (
	InvoiceUnPaid  InvoiceStatus = "unpaid"
	InvoicePaid    InvoiceStatus = "paid"
	InvoiceOverdue InvoiceStatus = "overdue"
)

type InvoiceModel struct {
	ID            uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	TransactionID uuid.UUID `gorm:"type:char(36);not null" json:"transaction_id"`
	InvoiceNumber string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"invoice_number"`
	DueDate       time.Time `gorm:"type:date;not null" json:"due_date"`
	Status        string    `gorm:"type:varchar(20);not null;default:unpaid" json:"status"`
	Notes         string    `gorm:"type:text" json:"notes"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Transaction *TransactionModel `gorm:"foreignKey:TransactionID" json:"transaction,omitempty"`
	Payments    []PaymentModel    `gorm:"foreignKey:InvoiceID" json:"payments,omitempty"`
}

func (InvoiceModel) TableName() string {
	return "invoices"
}

func (i *InvoiceModel) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}

	return
}
