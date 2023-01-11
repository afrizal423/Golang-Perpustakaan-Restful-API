package buku_request

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/oklog/ulid/v2"
)

type Jenis_Buku_Request struct {
	// gorm.Model
	JenisBuku string `json:"jenis_buku"  validate:"required,min=3"`
	Deskripsi string `json:"deskripsi" validate:"required,min=3"`
}

func (request *Jenis_Buku_Request) CreateJenisBuku() *models.Jenis_Buku {
	var jenbukRequest models.Jenis_Buku
	jenbukRequest.IDJenis = ulid.Make().String()
	jenbukRequest.JenisBuku = request.JenisBuku
	jenbukRequest.Deskripsi = request.Deskripsi
	return &jenbukRequest
}
