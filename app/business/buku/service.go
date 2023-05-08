package buku

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
)

type bukuService struct {
	repository IBukuRepository
}

func NewBukuService(repository IBukuRepository) IBukuService {
	return &bukuService{
		repository,
	}
}

func (s *bukuService) GetAllJenisBuku() ([]models.Jenis_Buku, error) {
	data, err := s.repository.GetAllJenisBuku()
	if err != nil {
		return []models.Jenis_Buku{}, nil
	}
	return data, nil
}

func (s *bukuService) FindJenisBuku(c string) ([]models.Jenis_Buku, error) {
	data, err := s.repository.CariJenisBuku(c)
	if err != nil {
		return []models.Jenis_Buku{}, nil
	}
	return data, nil
}

func (s *bukuService) GetJenisBukuById(id string) (*models.Jenis_Buku, error) {
	JenisBukuById, err := s.repository.GetJenisBukuById(id)
	if err != nil {
		return nil, err
	}
	return JenisBukuById, nil
}

func (s *bukuService) CreateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error) {
	var jb models.Jenis_Buku
	data, err := s.repository.CreateJenisBuku(data)
	if err != nil {
		return jb, err
	}
	return data, nil
}

func (s *bukuService) UpdateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error) {
	if s.repository.HitungDataJenisBuku(data.IDJenis) == 0 {
		return models.Jenis_Buku{}, errors.New("data kosong")
	}
	data, err := s.repository.UpdateJenisBuku(data)
	if err != nil {
		return models.Jenis_Buku{}, err
	}
	return data, nil
}

func (s *bukuService) HapusJenisBuku(id string) error {
	if s.repository.HitungDataJenisBuku(id) == 0 {
		return errors.New("data kosong")
	}
	if err := s.repository.DeleteJenisBuku(id); err != nil {
		return err
	}
	return nil
}
