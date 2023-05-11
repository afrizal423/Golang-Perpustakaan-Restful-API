package buku

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type IBukuService interface {
	// jenis buku

	GetAllJenisBuku() ([]models.Jenis_Buku, error)
	FindJenisBuku(c string) ([]models.Jenis_Buku, error)
	GetJenisBukuById(id string) (*models.Jenis_Buku, error)
	CreateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error)
	UpdateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error)
	HapusJenisBuku(id string) error

	// penerbit buku

	GetAllPenerbitBuku() ([]models.Penerbit_Buku, error)
	FindPenerbitBuku(c string) ([]models.Penerbit_Buku, error)
	GetPenerbitBukuById(id string) (*models.Penerbit_Buku, error)
	CreatePenerbitBuku(data models.Penerbit_Buku) (models.Penerbit_Buku, error)
	UpdatePenerbitBuku(data models.Penerbit_Buku) (models.Penerbit_Buku, error)
	HapusPenerbitBuku(id string) error
}

//go:generate mockery --name IBukuRepository
type IBukuRepository interface {
	// jenis buku

	GetAllJenisBuku() ([]models.Jenis_Buku, error)
	CariJenisBuku(c string) ([]models.Jenis_Buku, error)
	GetJenisBukuById(id string) (*models.Jenis_Buku, error)
	CreateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error)
	UpdateJenisBuku(data models.Jenis_Buku) (models.Jenis_Buku, error)
	DeleteJenisBuku(id string) error
	HitungDataJenisBuku(id string) int64

	// penerbit buku

	GetAllPenerbitBuku() ([]models.Penerbit_Buku, error)
	CariPenerbitBuku(c string) ([]models.Penerbit_Buku, error)
	GetPenerbitBukuById(id string) (*models.Penerbit_Buku, error)
	CreatePenerbitBuku(data models.Penerbit_Buku) (models.Penerbit_Buku, error)
	UpdatePenerbitBuku(data models.Penerbit_Buku) (models.Penerbit_Buku, error)
	DeletePenerbitBuku(id string) error
	HitungDataPenerbitBuku(id string) int64
}
