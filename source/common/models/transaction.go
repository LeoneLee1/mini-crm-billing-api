package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionStatus string

const (
	TransactionDraft     TransactionStatus = "draft"
	TransactionConfirmed TransactionStatus = "confirmed"
	TransactionCancelled TransactionStatus = "cancelled"
)

type TransactionModel struct {
	ID         uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	CustomerID uuid.UUID `gorm:"type:char(36)" json:"customer_id"`
	CreatedBy  uuid.UUID `gorm:"type:char(36)" json:"created_by"`
	Status     string    `gorm:"type:varchar(20);not null;default:draft" json:"status"`
	SubTotal   float64   `gorm:"type:decimal(15,2);not null;default:0" json:"sub_total"`
	Discount   float64   `gorm:"type:decimal(15,2);default:0" json:"discount"`
	Tax        float64   `gorm:"type:decimal(15,2);default:0" json:"tax"`
	Total      float64   `gorm:"type:decimal(15,2);not null;default:0" json:"total"`
	Notes      string    `gorm:"type:text" json:"notes"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Customer *CustomerModel         `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Creator  *UserModel             `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	Items    []TransactionItemModel `gorm:"foreignKey:TransactionID" json:"items,omitempty"`
	Invoice  *InvoiceModel          `gorm:"foreignKey:TransactionID" json:"invoice,omitempty"`
}

func (TransactionModel) TableName() string {
	return "transactions"
}

func (t *TransactionModel) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return

}

type TransactionItemModel struct {
	ID            uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	TransactionID uuid.UUID `gorm:"type:char(36)" json:"transaction_id"`
	ProductID     uuid.UUID `gorm:"type:char(36)" json:"product_id"`
	Quantity      float64   `gorm:"type:decimal(10,2);not null" json:"quantity"`
	UnitPrice     float64   `gorm:"type:decimal(15,2)" json:"unit_price"`
	SubTotal      float64   `gorm:"type:decimal(15,2);not null" json:"sub_total"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Product *ProductModel `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (TransactionItemModel) TableName() string {
	return "transaction_items"
}

func (ti *TransactionItemModel) BeforeCreate(tx *gorm.DB) (err error) {
	if ti.ID == uuid.Nil {
		ti.ID = uuid.New()
	}

	return
}
