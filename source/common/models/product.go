package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductModel struct {
	ID          uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(15,2);not null" json:"price"`
	Unit        string    `gorm:"type:varchar(255);not null" json:"unit"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedBy   uuid.UUID `gorm:"type:char(36)" json:"created_by"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Creator *UserModel `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (ProductModel) TableName() string {
	return "products"
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	return
}
