package buku

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type bukuService struct {
	repository IBukuRepository
}

func NewBukuService(repository IBukuRepository) IBukuRepository {
	return &bukuService{
		repository,
	}
}

func (s *bukuService) CreateJenisBuku(data models.Jenis_Buku) error {
	if err := s.repository.CreateJenisBuku(data); err != nil {
		return err
	}
	return nil
}
