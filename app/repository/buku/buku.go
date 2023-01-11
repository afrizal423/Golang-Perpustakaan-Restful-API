package buku

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (q *UserRepository) CreateJenisBuku(diary models.Jenis_Buku) error {
	if err := q.db.Save(&diary).Error; err != nil {
		return err
	}
	return nil
}
