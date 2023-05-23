package buku_response

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type ResponseSingleDataPenerbit struct {
	Error  bool                  `json:"error"`
	Status string                `json:"status"`
	Data   *models.Penerbit_Buku `json:"data"`
}

type ResponseManyDataPenerbit struct {
	Error  bool                   `json:"error"`
	Status string                 `json:"status"`
	Data   []models.Penerbit_Buku `json:"data"`
}

func GetPenerbitBukuByIdResponse(data *models.Penerbit_Buku, msg string, err bool) ResponseSingleDataPenerbit {
	var dt ResponseSingleDataPenerbit
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}

func GetAllPenerbitBukuResponse(data []models.Penerbit_Buku, msg string, err bool) ResponseManyDataPenerbit {
	var dt ResponseManyDataPenerbit
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}
