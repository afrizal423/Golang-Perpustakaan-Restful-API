package buku

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type IBukuService interface {
	GetJenisBukuById(id string) (*models.Jenis_Buku, error)
	CreateJenisBuku(data models.Jenis_Buku) error
	UpdateJenisBuku(data models.Jenis_Buku) error
	HapusJenisBuku(id string) error
}

type IBukuRepository interface {
	GetJenisBukuById(id string) (*models.Jenis_Buku, error)
	CreateJenisBuku(data models.Jenis_Buku) error
	UpdateJenisBuku(data models.Jenis_Buku) error
	DeleteJenisBuku(id string) error
	HitungDataJenisBuku(id string) int64
}
