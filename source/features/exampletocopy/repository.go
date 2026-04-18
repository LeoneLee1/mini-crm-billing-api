package exampletocopy

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	LogUserAccess(ctx context.Context, userID uint, requesterIP string) error
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}