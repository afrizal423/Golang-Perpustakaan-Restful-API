package buku

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
)

func (s *BukuService) GetAllJenisBuku() ([]models.Jenis_Buku, error) {
	data, err := s.Repository.GetAllJenisBuku()
	if err != nil {
		return []models.Jenis_Buku{}, nil
	}
	return data, nil
}

func (s *BukuService) FindJenisBuku(c string) ([]models.Jenis_Buku, error) {
	data, err := s.Repository.CariJenisBuku(c)
	if err != nil {
		return []models.Jenis_Buku{}, nil
	}
	return data, nil
}

func (s *BukuService) GetJenisBukuById(id string) (*models.Jenis_Buku, error) {
	JenisBukuById, err := s.Repository.GetJenisBukuById(id)
	if err != nil {
		return nil, err
	}
	return JenisBukuById, nil
}

func (s *BukuService) CreateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error) {
	var jb models.Jenis_Buku
	data, err := s.Repository.CreateJenisBuku(data)
	if err != nil {
		return jb, err
	}
	return data, nil
}

func (s *BukuService) UpdateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error) {
	if s.Repository.HitungDataJenisBuku(data.IDJenis) == 0 {
		return models.Jenis_Buku{}, errors.New("data kosong")
	}
	data, err := s.Repository.UpdateJenisBuku(data)
	if err != nil {
		return models.Jenis_Buku{}, err
	}
	return data, nil
}

func (s *BukuService) HapusJenisBuku(id string) error {
	if s.Repository.HitungDataJenisBuku(id) == 0 {
		return errors.New("data kosong")
	}
	if err := s.Repository.DeleteJenisBuku(id); err != nil {
		return err
	}
	return nil
}
