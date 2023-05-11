package buku_request

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"
	"github.com/oklog/ulid/v2"
)

type Penerbit_Buku_Request struct {
	Deskripsi      string `json:"deskripsi" validate:"required"`
	PenerbitBuku   string `json:"penerbit_buku" validate:"required"`
	AlamatPenerbit string ` json:"alamat_penerbit" validate:"required"`
	TelpPenerbit   string `json:"telp_penerbit" validate:"required"`
	EmailPenerbit  string `json:"email_penerbit" validate:"required"`
}
type Penerbit_Buku_Request_Update struct {
	IDPenerbit     string `json:"id" validate:"required,min=26"`
	Deskripsi      string `json:"deskripsi" validate:"required"`
	PenerbitBuku   string `json:"penerbit_buku" validate:"required"`
	AlamatPenerbit string ` json:"alamat_penerbit" validate:"required"`
	TelpPenerbit   string `json:"telp_penerbit" validate:"required"`
	EmailPenerbit  string `json:"email_penerbit" validate:"required"`
}
type Penerbit_Buku_Request_Delete struct {
	IDPenerbit string `json:"id" validate:"required,min=26"`
}

func (request *Penerbit_Buku_Request) CreatePenerbitBuku() *models.Penerbit_Buku {
	var penbukRequest models.Penerbit_Buku
	penbukRequest.IDPenerbit = ulid.Make().String()
	penbukRequest.PenerbitBuku = request.PenerbitBuku
	penbukRequest.AlamatPenerbit = request.AlamatPenerbit
	penbukRequest.TelpPenerbit = request.TelpPenerbit
	penbukRequest.EmailPenerbit = request.EmailPenerbit
	penbukRequest.Deskripsi = request.Deskripsi
	return &penbukRequest
}

func (request *Penerbit_Buku_Request_Update) UpdatePenerbitBuku(idnya string) *models.Penerbit_Buku {
	var penbukRequest models.Penerbit_Buku
	penbukRequest.IDPenerbit = idnya
	penbukRequest.PenerbitBuku = request.PenerbitBuku
	penbukRequest.AlamatPenerbit = request.AlamatPenerbit
	penbukRequest.TelpPenerbit = request.TelpPenerbit
	penbukRequest.EmailPenerbit = request.EmailPenerbit
	penbukRequest.Deskripsi = request.Deskripsi
	return &penbukRequest
}
