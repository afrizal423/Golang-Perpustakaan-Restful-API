package buku_response

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type ResponseSingleDataPenulis struct {
	Error  bool                 `json:"error"`
	Status string               `json:"status"`
	Data   *models.Penulis_Buku `json:"data"`
}

type ResponseManyDataPenulis struct {
	Error  bool                  `json:"error"`
	Status string                `json:"status"`
	Data   []models.Penulis_Buku `json:"data"`
}

func GetPenulisBukuByIdResponse(data *models.Penulis_Buku, msg string, err bool) ResponseSingleDataPenulis {
	var dt ResponseSingleDataPenulis
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}

func GetAllPenulisBukuResponse(data []models.Penulis_Buku, msg string, err bool) ResponseManyDataPenulis {
	var dt ResponseManyDataPenulis
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}
