package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogActivityModel struct {
	ID         uuid.UUID  `gorm:"primaryKey;type:char(36)" json:"id"`
	UserID     uuid.UUID  `gorm:"index;type:char(36)" json:"user_id"`
	Action     string     `gorm:"type:text" json:"action"`
	DataBefore string     `gorm:"type:json" json:"data_before"`
	DataAfter  string     `gorm:"type:json" json:"data_after"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	User       *UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (LogActivityModel) TableName() string {
	return "log_activities"
}

func (l *LogActivityModel) BeforeCreate(tx *gorm.DB) (err error) {
	if l.DataBefore == "" {
		l.DataBefore = "{}"
	}
	if l.DataAfter == "" {
		l.DataAfter = "{}"
	}
	return
}
