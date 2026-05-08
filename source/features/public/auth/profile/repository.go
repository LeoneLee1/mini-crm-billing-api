package profile

import (
	"gorm.io/gorm"
)

type Repository interface {
}

type repositoryImpl struct {
	db *gorm.DB
}

func injectRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}
