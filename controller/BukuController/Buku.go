package BukuController

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config/helper"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models/BukuModels"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
	"github.com/biezhi/gorm-paginator/pagination"
)

func LihatSemuaBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.Buku{}
	err, querydb := buku.LihatSemuaBuku(config.Db)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	query := r.URL.Query()
	id, _ := strconv.ParseUint(query.Get("id"), 10, 32)
	// Lihat Detail jenis buku
	if id != 0 {
		bukus, err := buku.DetailBuku(config.Db, uint32(id))
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		// fmt.Println(bukus)
		responses.JSON(w, http.StatusOK, responses.Sukses(bukus))
		return
	}
	if query.Get("search") == "true" {
		cariBuku(w, r)
		return
	}
	var tmpbuku []BukuModels.Buku
	halaman, _ := strconv.ParseUint(query.Get("page"), 10, 32)
	paginator := pagination.Paging(&pagination.Param{
		DB:    querydb,
		Page:  int(halaman),
		Limit: 3,
		// OrderBy: []string{"id desc"},
	}, &tmpbuku)
	responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
}

func cariBuku(w http.ResponseWriter, r *http.Request) {
	buku := BukuModels.Buku{}
	query := r.URL.Query()
	if query.Get("q") != "" && query.Get("orderby") != "" {
		kata := helper.NonLatinChar(query.Get("q"))[0]
		orderby := helper.NonLatinChar(query.Get("orderby"))[0]
		_, data := buku.SearchBuku(config.Db, kata, orderby)
		halaman, _ := strconv.ParseUint(query.Get("page"), 10, 32)

		var bukus []BukuModels.Buku
		paginator := pagination.Paging(&pagination.Param{
			DB:    data,
			Page:  int(halaman),
			Limit: 3,
			// OrderBy: []string{"id desc"},
		}, &bukus)
		responses.JSON(w, http.StatusOK, responses.Sukses(paginator))
		fmt.Println(kata, orderby)
		return
	}
}
