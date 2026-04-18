package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenModel struct {
	ID           uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	UserID       uuid.UUID `gorm:"index;type:char(36)" json:"user_id"`
	RefreshToken string    `gorm:"type:text" json:"refresh_token"`
	ExpiredAt    time.Time `gorm:"index" json:"expired_at"`
	Revoked      bool      `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"user,omitempty"`
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}

func (r *RefreshTokenModel) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}

	return
}
