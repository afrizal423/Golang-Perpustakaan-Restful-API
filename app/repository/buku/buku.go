package buku

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}
