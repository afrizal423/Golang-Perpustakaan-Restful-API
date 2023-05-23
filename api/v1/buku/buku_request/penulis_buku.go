package buku_request

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/oklog/ulid/v2"
)

type Penulis_Buku_Request struct {
	PenulisBuku   string `validate:"required" json:"penulis_buku"`
	AlamatPenulis string `validate:"required" json:"alamat_penulis"`
	EmailPenulis  string `validate:"required" json:"email_penulis"`
	Deskripsi     string `validate:"required" json:"deskripsi"`
}

type Penulis_Buku_Request_Update struct {
	IDPenulis     string `validate:"required" json:"id"`
	PenulisBuku   string `validate:"required" json:"penulis_buku"`
	AlamatPenulis string `validate:"required" json:"alamat_penulis"`
	EmailPenulis  string `validate:"required" json:"email_penulis"`
	Deskripsi     string `validate:"required" json:"deskripsi"`
}

type Penulis_Buku_Request_Delete struct {
	IDPenulis string `validate:"required" json:"id"`
}

func (request *Penulis_Buku_Request) CreatePenulisBuku() *models.Penulis_Buku {
	var penbukRequest models.Penulis_Buku
	penbukRequest.IDPenulis = ulid.Make().String()
	penbukRequest.PenulisBuku = request.PenulisBuku
	penbukRequest.AlamatPenulis = request.AlamatPenulis
	penbukRequest.EmailPenulis = request.EmailPenulis
	penbukRequest.Deskripsi = request.Deskripsi
	return &penbukRequest
}

func (request *Penulis_Buku_Request_Update) UpdatePenulisBuku(idnya string) *models.Penulis_Buku {
	var penbukRequest models.Penulis_Buku
	penbukRequest.IDPenulis = idnya
	penbukRequest.PenulisBuku = request.PenulisBuku
	penbukRequest.AlamatPenulis = request.AlamatPenulis
	penbukRequest.EmailPenulis = request.EmailPenulis
	penbukRequest.Deskripsi = request.Deskripsi
	return &penbukRequest
}
