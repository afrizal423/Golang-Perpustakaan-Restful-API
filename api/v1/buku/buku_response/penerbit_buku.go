package buku_response

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/app/models"

type ResponseSingleData struct {
	Error  bool                  `json:"error"`
	Status string                `json:"status"`
	Data   *models.Penerbit_Buku `json:"data"`
}

type ResponseManyData struct {
	Error  bool                   `json:"error"`
	Status string                 `json:"status"`
	Data   []models.Penerbit_Buku `json:"data"`
}

func GetPenerbitBukuByIdResponse(data *models.Penerbit_Buku, msg string, err bool) ResponseSingleData {
	var dt ResponseSingleData
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}

func GetAllPenerbitBukuResponse(data []models.Penerbit_Buku, msg string, err bool) ResponseManyData {
	var dt ResponseManyData
	dt.Error = err
	dt.Status = msg
	dt.Data = data
	return dt
}
