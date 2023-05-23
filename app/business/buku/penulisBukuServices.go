package buku

import (
	"errors"
	"regexp"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
)

func (s *BukuService) GetAllPenulisBuku() ([]models.Penulis_Buku, error) {
	data, err := s.Repository.GetAllPenulisBuku()
	if err != nil {
		return []models.Penulis_Buku{}, nil
	}
	if len(data) == 0 {
		return []models.Penulis_Buku{}, errors.New("data kosong")
	}
	return data, nil
}

func (s *BukuService) FindPenulisBuku(c string) ([]models.Penulis_Buku, error) {
	data, err := s.Repository.CariPenulisBuku(c)
	if err != nil {
		return []models.Penulis_Buku{}, nil
	}
	if len(data) == 0 {
		return []models.Penulis_Buku{}, errors.New("data kosong")
	}
	return data, nil
}

func (s *BukuService) GetPenulisBukuById(id string) (*models.Penulis_Buku, error) {
	JenisBukuById, err := s.Repository.GetPenulisBukuById(id)
	if err != nil {
		return nil, err
	}
	if JenisBukuById.IDPenulis == "" {
		return &models.Penulis_Buku{}, errors.New("data kosong")
	}
	return JenisBukuById, nil
}

func (s *BukuService) CreatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error) {
	var jb models.Penulis_Buku

	// validasi inputan nama penulis buku
	if len(data.PenulisBuku) <= 5 || len(data.PenulisBuku) >= 255 {
		return models.Penulis_Buku{}, errors.New("inputan nama penulis buku minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan alamat penulis buku
	if len(data.AlamatPenulis) <= 5 || len(data.AlamatPenulis) >= 255 {
		return models.Penulis_Buku{}, errors.New("inputan alamat penulis buku minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan email Penulis
	match, _ := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", data.EmailPenulis)

	if !match {
		return models.Penulis_Buku{}, errors.New("bukan format email")
	}

	// validasi inputan deskripsi penulis buku
	if len(data.PenulisBuku) >= 255 {
		return models.Penulis_Buku{}, errors.New("inputan deskripsi penulis buku maksimal 255")
	}
	// end validation

	data, err := s.Repository.CreatePenulisBuku(data)
	if err != nil {
		return jb, err
	}
	return data, nil
}

func (s *BukuService) UpdatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error) {
	var jb models.Penulis_Buku
	if s.Repository.HitungDataPenulisBuku(data.IDPenulis) == 0 {
		return models.Penulis_Buku{}, errors.New("data kosong")
	}
	// validasi inputan nama penulis buku
	if len(data.PenulisBuku) <= 5 || len(data.PenulisBuku) >= 255 {
		return models.Penulis_Buku{}, errors.New("inputan nama penulis buku minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan alamat penulis buku
	if len(data.AlamatPenulis) <= 5 || len(data.AlamatPenulis) >= 255 {
		return models.Penulis_Buku{}, errors.New("inputan alamat penulis buku minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan email Penulis
	match, _ := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", data.EmailPenulis)

	if !match {
		return models.Penulis_Buku{}, errors.New("bukan format email")
	}

	// validasi inputan deskripsi penulis buku
	if len(data.PenulisBuku) >= 255 {
		return models.Penulis_Buku{}, errors.New("inputan deskripsi penulis buku maksimal 255")
	}
	// end validation

	data, err := s.Repository.UpdatePenulisBuku(data)
	if err != nil {
		return jb, err
	}
	return data, nil
}

func (s *BukuService) HapusPenulisBuku(id string) error {
	if s.Repository.HitungDataPenulisBuku(id) == 0 {
		return errors.New("data kosong")
	}
	if err := s.Repository.DeletePenulisBuku(id); err != nil {
		return err
	}
	return nil
}
