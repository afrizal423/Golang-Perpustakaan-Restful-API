package user

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (q *UserRepository) GetDataByUsername(username string) (*models.Pegawai, error) {
	var data models.Pegawai
	if err := q.db.Where("username = ?", username).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
