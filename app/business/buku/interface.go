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

	// penulis buku

	GetAllPenulisBuku() ([]models.Penulis_Buku, error)
	FindPenulisBuku(c string) ([]models.Penulis_Buku, error)
	GetPenulisBukuById(id string) (*models.Penulis_Buku, error)
	CreatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error)
	UpdatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error)
	HapusPenulisBuku(id string) error
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

	// penulis buku

	CreatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error)
	UpdatePenulisBuku(data models.Penulis_Buku) (models.Penulis_Buku, error)
	DeletePenulisBuku(id string) error
	HitungDataPenulisBuku(id string) int64
	GetAllPenulisBuku() ([]models.Penulis_Buku, error)
	CariPenulisBuku(c string) ([]models.Penulis_Buku, error)
	GetPenulisBukuById(id string) (*models.Penulis_Buku, error)
}
