package buku

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
)

func (q *BukuRepository) CreatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error) {
	var jb models.Penulis_Buku
	if err := q.db.Save(&data).Error; err != nil {
		return jb, err
	}
	return data, nil
}

func (q *BukuRepository) UpdatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error) {
	if err := q.db.Where("id_penulis = ?", &data.IDPenulis).Updates(&data).Error; err != nil {
		return models.Penulis_Buku{}, err
	}
	return data, nil
}

func (q *BukuRepository) HitungDataPenulisBuku(id string) int64 {
	var count int64
	if err := q.db.Model(&models.Penulis_Buku{}).Where("id_penulis = ?", id).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return count
}

func (q *BukuRepository) DeletePenulisBuku(id string) error {
	var data models.Penulis_Buku
	if err := q.db.Where("id_penulis = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *BukuRepository) GetAllPenulisBuku() ([]models.Penulis_Buku, error) {
	var data, result []models.Penulis_Buku
	if err := q.db.Find(&data).Error; err != nil {
		return nil, err
	}
	// for _, v := range data {
	result = append(result, data...)
	// }
	return result, nil
}

func (q *BukuRepository) CariPenulisBuku(c string) ([]models.Penulis_Buku, error) {
	var data, result []models.Penulis_Buku
	if err := q.db.Where("penulis_buku LIKE ? OR email_penulis LIKE ? OR deskripsi LIKE ?", "%"+c+"%", "%"+c+"%", "%"+c+"%").Find(&data).Error; err != nil {
		return nil, err
	}
	result = append(result, data...)
	return result, nil
}

func (q *BukuRepository) GetPenulisBukuById(id string) (*models.Penulis_Buku, error) {
	var data models.Penulis_Buku
	if err := q.db.First(&data, "id_penulis = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
