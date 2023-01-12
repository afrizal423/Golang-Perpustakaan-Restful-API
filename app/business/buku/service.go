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

func (s *bukuService) GetJenisBukuById(id string) (*models.Jenis_Buku, error) {
	JenisBukuById, err := s.repository.GetJenisBukuById(id)
	if err != nil {
		return nil, err
	}
	return JenisBukuById, nil
}

func (s *bukuService) CreateJenisBuku(data models.Jenis_Buku) error {
	if err := s.repository.CreateJenisBuku(data); err != nil {
		return err
	}
	return nil
}

func (s *bukuService) UpdateJenisBuku(data models.Jenis_Buku) error {
	if s.repository.HitungDataJenisBuku(data.IDJenis) == 0 {
		return errors.New("data kosong")
	}
	if err := s.repository.UpdateJenisBuku(data); err != nil {
		return err
	}
	return nil
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
