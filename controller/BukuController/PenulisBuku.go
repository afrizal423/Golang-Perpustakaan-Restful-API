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

func LihatPenulisBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.PenulisBuku{}
	bukus, err := buku.LihatPenulisBuku(config.Db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	// Lihat Detail jenis buku
	if id != 0 {
		bukus, err := buku.DetailPenulisBuku(config.Db, uint32(id))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		// fmt.Println(bukus)
		responses.JSON(w, http.StatusOK, responses.Sukses(bukus))
		return
	}
	halaman, _ := strconv.ParseUint(query.Get("page"), 10, 32)
	paginator := pagination.Paging(&pagination.Param{
		DB:    config.Db,
		Page:  int(halaman),
		Limit: 3,
		// OrderBy: []string{"id desc"},
	}, &bukus)
	responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
}

func TambahPenulisBuku(w http.ResponseWriter, r *http.Request) {
	var penbuk BukuModels.PenulisBuku
	err := json.NewDecoder(r.Body).Decode(&penbuk)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	err = penbuk.Validasi()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tmbhData, err := penbuk.TambahPenulisBuku(config.Db)
	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, tmbhData.IDPenulis))
	responses.JSON(w, http.StatusCreated, tmbhData)
}

func HapusPenulisBuku(w http.ResponseWriter, r *http.Request) {
	var jenbuk BukuModels.PenulisBuku
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	if id == 0 {
		responses.ERROR(w, http.StatusInternalServerError, fmt.Errorf("URL Invalid"))
		return
	}
	_, err := jenbuk.HapusPenulisBuku(config.Db, uint32(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusOK, responses.Sukses(fmt.Sprintf("Data dengan id %d sukses terhapus ", id)))
}