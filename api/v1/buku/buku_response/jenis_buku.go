package buku_response

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

var (
	ErrInvalidFormatJson = "Data json yang diberikan salah!. Tidak sesuai format"
)

type GetJenisBukuById struct {
	Status string              `json:"status"`
	Data   dataDetailJenisBuku `json:"data"`
}

type dataDetailJenisBuku struct {
	IDJenis   string `json:"id"`
	JenisBuku string `json:"jenis_buku"`
	Deskripsi string `json:"deskripsi"`
}

func GetJenisBukuByIdResponse(data *models.Jenis_Buku) GetJenisBukuById {
	var dt GetJenisBukuById
	dt.Status = "sukses"
	dt.Data.IDJenis = data.IDJenis
	dt.Data.JenisBuku = data.JenisBuku
	dt.Data.Deskripsi = data.Deskripsi
	return dt
}
