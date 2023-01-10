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

func (q *UserRepository) Login(email string, password string) (*models.Anggota, error) {
	var data models.Anggota
	if err := q.db.Where("email = ?", email).Where("password = ?", password).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
