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
	// diaryById, err := s.repository.GetDiaryById(id)
	// if err != nil {
	// 	return nil, err
	// }
	return nil
}
