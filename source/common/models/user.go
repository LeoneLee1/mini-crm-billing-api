package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleAccess string

const (
	RoleStaff RoleAccess = "staff"
	RoleAdmin RoleAccess = "admin"
)

type UserModel struct {
	ID       uuid.UUID  `gorm:"primaryKey;type:char(36)" json:"id"`
	Name     string     `gorm:"type:varchar(255)" json:"name"`
	Email    string     `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password string     `gorm:"type:varchar(255)" json:"-"`
	Role     RoleAccess `gorm:"type:varchar(100)" json:"role"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	return
}
