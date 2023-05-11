package buku_response

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

var (
	ErrInvalidFormatJson = "Data json yang diberikan salah!. Tidak sesuai format"
)

type GetJenisBukuById struct {
	Status string             `json:"status"`
	Data   *models.Jenis_Buku `json:"data"`
}

func GetJenisBukuByIdResponse(data *models.Jenis_Buku) GetJenisBukuById {
	var dt GetJenisBukuById
	dt.Status = "Success get data buku"
	dt.Data = data
	return dt
}
