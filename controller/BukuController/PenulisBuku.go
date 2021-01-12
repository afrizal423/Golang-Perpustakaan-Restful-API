package BukuController

import (
	"net/http"
	"strconv"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/models/BukuModels"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/responses"
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
