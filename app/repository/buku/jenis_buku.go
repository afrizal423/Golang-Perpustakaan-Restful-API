package buku

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
)

func (q *BukuRepository) GetAllJenisBuku() ([]models.Jenis_Buku, error) {
	var data, result []models.Jenis_Buku
	if err := q.db.Find(&data).Error; err != nil {
		return nil, err
	}
	// for _, v := range data {
	result = append(result, data...)
	// }
	return result, nil
}

func (q *BukuRepository) CariJenisBuku(c string) ([]models.Jenis_Buku, error) {
	var data, result []models.Jenis_Buku
	if err := q.db.Where("jenis_buku LIKE ? OR deskripsi LIKE ?", "%"+c+"%", "%"+c+"%").Find(&data).Error; err != nil {
		return nil, err
	}
	result = append(result, data...)
	return result, nil
}

func (q *BukuRepository) GetJenisBukuById(id string) (*models.Jenis_Buku, error) {
	var data models.Jenis_Buku
	if err := q.db.First(&data, "id_jenis = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (q *BukuRepository) CreateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error) {
	var jb models.Jenis_Buku
	if err := q.db.Save(&data).Error; err != nil {
		return jb, err
	}
	return data, nil
}

func (q *BukuRepository) UpdateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error) {
	if err := q.db.Where("id_jenis = ?", &data.IDJenis).Updates(&data).Error; err != nil {
		return models.Jenis_Buku{}, err
	}
	return data, nil
}

func (q *BukuRepository) DeleteJenisBuku(id string) error {
	var data models.Jenis_Buku
	if err := q.db.Where("id_jenis = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *BukuRepository) HitungDataJenisBuku(id string) int64 {
	var count int64
	if err := q.db.Model(&models.Jenis_Buku{}).Where("id_jenis = ?", id).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return count
}
