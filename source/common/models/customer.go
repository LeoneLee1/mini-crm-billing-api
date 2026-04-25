package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerStatus string

const (
	CustomerStatusActive   CustomerStatus = "active"   // can be selected for create transaction
	CustomerStatusInActive CustomerStatus = "inactive" // does not appear in the customer selection dropdow when creating a new transaction
)

type CustomerModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	Status    string    `gorm:"type:varchar(20);not null;default:active" json:"status"`
	CreatedBy uuid.UUID `gorm:"type:char(36)" json:"created_by"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // rules: if customer have a data transaction, this customer can't be deleted

	Creator *UserModel `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (CustomerModel) TableName() string {
	return "customers"
}

func (c *CustomerModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return
}
