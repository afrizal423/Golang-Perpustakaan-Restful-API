package buku

import (
	"log"

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

func (q *UserRepository) GetAllJenisBuku() ([]models.Jenis_Buku, error) {
	var data, result []models.Jenis_Buku
	if err := q.db.Find(&data).Error; err != nil {
		return nil, err
	}
	for _, v := range data {
		result = append(result, v)
	}
	return result, nil
}

func (q *UserRepository) GetJenisBukuById(id string) (*models.Jenis_Buku, error) {
	var data models.Jenis_Buku
	if err := q.db.Debug().First(&data, "id_jenis = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (q *UserRepository) CreateJenisBuku(data models.Jenis_Buku) error {
	if err := q.db.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *UserRepository) UpdateJenisBuku(data models.Jenis_Buku) error {
	if err := q.db.Where("id_jenis = ?", &data.IDJenis).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *UserRepository) DeleteJenisBuku(id string) error {
	var data models.Jenis_Buku
	if err := q.db.Where("id_jenis = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *UserRepository) HitungDataJenisBuku(id string) int64 {
	var count int64
	if err := q.db.Model(&models.Jenis_Buku{}).Where("id_jenis = ?", id).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return count
}
