package BukuController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models/BukuModels"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses/formaterror"
	"github.com/biezhi/gorm-paginator/pagination"
)

func LihatPenerbitBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.PenerbitBuku{}
	bukus, err := buku.LihatPenerbitBuku(config.Db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	query := r.URL.Query()
	halaman, _ := strconv.ParseUint(query.Get("page"), 10, 32)
	paginator := pagination.Paging(&pagination.Param{
		DB:    config.Db,
		Page:  int(halaman),
		Limit: 3,
		// OrderBy: []string{"id desc"},
	}, &bukus)
	responses.JSON(w, http.StatusOK, responses.Sukses(paginator))

}

func TambahPenerbitBuku(w http.ResponseWriter, r *http.Request) {
	var pubbuk BukuModels.PenerbitBuku
	err := json.NewDecoder(r.Body).Decode(&pubbuk)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = pubbuk.Validasi()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tmbhData, err := pubbuk.TambahPenerbitBuku(config.Db)
	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, tmbhData.IDPenerbit))
	responses.JSON(w, http.StatusCreated, tmbhData)
}
